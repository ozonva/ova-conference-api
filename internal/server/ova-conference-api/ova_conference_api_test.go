package ova_conference_api_test

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"ova-conference-api/internal/domain"
	ovaconferenceapi "ova-conference-api/internal/server/ova-conference-api"
	"ova-conference-api/internal/utils/mocks"
	conf "ova-conference-api/pkg/api/github.com/ozonva/ova-conference-api/pkg/ova-conference-api"
	"time"
)

var _ = Describe("Server test", func() {
	var (
		mockCtrl *gomock.Controller
		mockRepo *mocks.MockRepo
		api      conf.ConferencesServer
		ctx      context.Context
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(mockCtrl)
		ctx = context.Background()
		api = ovaconferenceapi.NewServer(mockRepo)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("CreateEntity", func() {
		It("success saving entity", func() {
			entity := domain.NewConference("test", &domain.EventTime{Time: time.Now()})
			mockRepo.EXPECT().AddEntity(ctx, gomock.Any()).Times(1).Return(entity, nil)
			_, err := api.CreateConference(ctx, ovaconferenceapi.ToConferenceRequest(entity))
			fmt.Println(err)

			Expect(err).ShouldNot(HaveOccurred())
		})

		It("DescribeEntity", func() {
			entity := domain.NewConference("test", &domain.EventTime{Time: time.Now()})
			entity.Id = 1
			mockRepo.EXPECT().DescribeEntity(ctx, int64(1)).Return(entity, nil).Times(1)

			res, err := api.DescribeConference(ctx, &conf.EntityConferenceRequest{
				Id: 1,
			})
			Expect(err).ShouldNot(HaveOccurred())

			Expect(res.Id).To(Equal(entity.Id))
			Expect(res.Name).To(Equal(entity.Name))
			Expect(res.SpeakerCount).To(Equal(entity.SpeakerCount))
			Expect(res.ParticipantCount).To(Equal(entity.ParticipantCount))
		})

		It("List invalid parameters", func() {
			_, err := api.ListConference(ctx, &conf.ListConferenceRequest{
				Limit:  int64(-1),
				Offset: int64(2),
			})
			Expect(err).NotTo(BeNil())
		})

		It("List", func() {
			entity := domain.MakeConference("test", &domain.EventTime{Time: time.Now()}, 1, 2)
			someItems := []domain.Conference{entity, entity, entity}
			mockRepo.EXPECT().ListEntities(ctx, int64(3), int64(0)).Return(someItems, nil).Times(1)

			res, err := api.ListConference(ctx, &conf.ListConferenceRequest{
				Limit:  int64(3),
				Offset: int64(0),
			})
			Expect(err).ShouldNot(HaveOccurred())

			Expect(len(res.Items)).To(Equal(len(someItems)))
		})

		It("RemoveEntity", func() {
			mockRepo.EXPECT().DeleteEntity(ctx, int64(1)).Return(nil).Times(1)

			_, err := api.RemoveConference(ctx, &conf.EntityConferenceRequest{
				Id: int64(1),
			})
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
