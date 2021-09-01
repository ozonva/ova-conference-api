package ova_conference_api

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"ova-conference-api/internal/domain"
	conf "ova-conference-api/pkg/api/github.com/ozonva/ova-conference-api/pkg/ova-conference-api"
)

func ToConferenceResponse(conference *domain.Conference) *conf.ConferenceResponse {
	return &conf.ConferenceResponse{
		Id:               conference.Id,
		Name:             conference.Name,
		EventTime:        timestamppb.New(conference.EventTime.Time),
		ParticipantCount: conference.ParticipantCount,
		SpeakerCount:     conference.SpeakerCount,
	}
}

func ToConferenceDomain(request *conf.CreateConferenceRequest) domain.Conference {
	return domain.MakeConference(request.Name, &domain.EventTime{Time: request.EventTime.AsTime()}, request.ParticipantCount, request.SpeakerCount)
}

func ToConferenceRequest(conference *domain.Conference) *conf.CreateConferenceRequest {
	return &conf.CreateConferenceRequest{
		Name:             conference.Name,
		EventTime:        timestamppb.New(conference.EventTime.Time),
		ParticipantCount: conference.ParticipantCount,
		SpeakerCount:     conference.SpeakerCount,
	}
}
