package flusher_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"ova-conference-api/internal/domain"
	"ova-conference-api/internal/utils/flusher"
	"ova-conference-api/internal/utils/mocks"
	"time"
)

var _ = Describe("Flusher", func() {
	var (
		mockCtrl    *gomock.Controller
		mockRepo    *mocks.MockRepo
		testFlusher flusher.Flusher
		ctx         context.Context
		conferences = []domain.Conference{
			domain.MakeConference("test", &domain.EventTime{Time: time.Now()}, 1, 2),
			domain.MakeConference("test2", &domain.EventTime{Time: time.Now()}, 1, 2),
			domain.MakeConference("test3", &domain.EventTime{Time: time.Now()}, 1, 2),
			domain.MakeConference("test4", &domain.EventTime{Time: time.Now()}, 1, 2),
		}
	)
	AfterEach(func() {
		mockCtrl.Finish()
	})
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(mockCtrl)
		testFlusher = flusher.NewFlusher(3, mockRepo)
	})
	Describe("All data are written", func() {
		It("Write one element", func() {
			oneItem := conferences[:1]
			mockRepo.EXPECT().AddEntities(ctx, oneItem, nil).Return(nil)
			Expect(testFlusher.Flush(ctx, oneItem)).To(BeNil())
		})

		It("Write less then chunkSize", func() {
			twoItems := conferences[:2]
			mockRepo.EXPECT().AddEntities(ctx, twoItems, nil).Return(nil)
			Expect(testFlusher.Flush(ctx, twoItems)).To(BeNil())
		})

		It("Write more then chunkSize", func() {
			someItems := conferences[:3]
			mockRepo.EXPECT().AddEntities(ctx, someItems, nil).Return(nil).AnyTimes()
			Expect(testFlusher.Flush(ctx, someItems)).To(BeNil())
		})
	})

	Describe("data has problems with saving", func() {
		It("All data failed - should return all elements", func() {
			err := errors.New("some error from repo")
			gomock.InOrder(
				mockRepo.EXPECT().AddEntities(ctx, conferences[:2], nil).Return(err).Times(1),
				mockRepo.EXPECT().AddEntities(ctx, conferences[2:], nil).Return(err).Times(1),
			)
			testFlusher = flusher.NewFlusher(2, mockRepo)
			Expect(testFlusher.Flush(ctx, conferences)).To(Equal(conferences))
		})

		It("Some data failed - should return something", func() {
			gomock.InOrder(
				mockRepo.EXPECT().AddEntities(ctx, conferences[:2], nil).Return(nil).Times(1),
				mockRepo.EXPECT().AddEntities(ctx, conferences[2:], nil).Return(errors.New("some error from repo")).Times(1),
			)
			testFlusher = flusher.NewFlusher(2, mockRepo)
			result := testFlusher.Flush(ctx, conferences)
			Expect(len(result)).Should(BeNumerically(">=", 1))
		})
	})
})
