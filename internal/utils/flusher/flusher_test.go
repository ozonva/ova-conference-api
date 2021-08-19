package flusher_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"ova-conference-api/internal/domain"
	"ova-conference-api/internal/utils/flusher"
	"ova-conference-api/internal/utils/repo/mocks"
	"time"
)

var _ = Describe("Flusher", func() {
	var (
		mockCtrl    *gomock.Controller
		mockRepo    *mocks.MockRepo
		testFlusher flusher.Flusher
		conferences = []domain.Conference{
			domain.MakeConference(1, "test", &domain.EventTime{Time: time.Now()}),
			domain.MakeConference(2, "test2", &domain.EventTime{Time: time.Now()}),
			domain.MakeConference(3, "test3", &domain.EventTime{Time: time.Now()}),
			domain.MakeConference(4, "test4", &domain.EventTime{Time: time.Now()}),
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
			mockRepo.EXPECT().AddEntities(oneItem).Return(nil)
			Expect(testFlusher.Flush(oneItem)).To(BeNil())
		})

		It("Write less then chunkSize", func() {
			twoItems := conferences[:2]
			mockRepo.EXPECT().AddEntities(twoItems).Return(nil)
			Expect(testFlusher.Flush(twoItems)).To(BeNil())
		})

		It("Write more then chunkSize", func() {
			someItems := conferences[:3]
			mockRepo.EXPECT().AddEntities(someItems).Return(nil).AnyTimes()
			Expect(testFlusher.Flush(someItems)).To(BeNil())
		})
	})

	Describe("data has problems with saving", func() {
		It("All data failed - should return all elements", func() {
			err := errors.New("some error from repo")
			gomock.InOrder(
				mockRepo.EXPECT().AddEntities(conferences[:2]).Return(err).Times(1),
				mockRepo.EXPECT().AddEntities(conferences[2:]).Return(err).Times(1),
			)
			testFlusher = flusher.NewFlusher(2, mockRepo)
			Expect(testFlusher.Flush(conferences)).To(Equal(conferences))
		})

		It("Some data failed - should return something", func() {
			gomock.InOrder(
				mockRepo.EXPECT().AddEntities(conferences[:2]).Return(nil).Times(1),
				mockRepo.EXPECT().AddEntities(conferences[2:]).Return(errors.New("some error from repo")).Times(1),
			)
			testFlusher = flusher.NewFlusher(2, mockRepo)
			result := testFlusher.Flush(conferences)
			Expect(len(result)).Should(BeNumerically(">=", 1))
		})
	})
})
