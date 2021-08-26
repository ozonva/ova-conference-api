package ova_conference_api

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math/rand"
	conf "ova-conference-api/pkg/api/github.com/ozonva/ova-conference-api/pkg/ova-conference-api"
	"time"
)

func getDummyConference() *conf.ConferenceResponse {
	rand.Seed(time.Now().UnixNano())
	return &conf.ConferenceResponse{

		Id:               &conf.UUID{Value: uuid.New().String()},
		Userid:           rand.Uint64(),
		Name:             "dummy_conf",
		EventTime:        timestamppb.New(time.Now()),
		ParticipantCount: rand.Int31(),
		SpeakerCount:     rand.Int31(),
	}
}
