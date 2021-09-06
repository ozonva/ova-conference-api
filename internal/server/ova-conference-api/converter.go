package ova_conference_api

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"ova-conference-api/internal/domain"
	conf "ova-conference-api/pkg/api/github.com/ozonva/ova-conference-api/pkg/ova-conference-api"
)

func FromDomainToConferenceEntity(conference *domain.Conference) *conf.ConferenceEntity {
	return &conf.ConferenceEntity{
		Id:               conference.Id,
		Name:             conference.Name,
		EventTime:        timestamppb.New(conference.EventTime.Time),
		ParticipantCount: conference.ParticipantCount,
		SpeakerCount:     conference.SpeakerCount,
	}
}

func FromRequestToDomain(request *conf.CreateConferenceRequest) domain.Conference {
	return domain.MakeConference(request.Name, &domain.EventTime{Time: request.EventTime.AsTime()}, request.ParticipantCount, request.SpeakerCount)
}

func FromConferenceEntityToDomain(request *conf.ConferenceEntity) domain.Conference {
	entity := domain.MakeConference(request.Name, &domain.EventTime{Time: request.EventTime.AsTime()}, request.ParticipantCount, request.SpeakerCount)
	entity.Id = request.Id
	return entity
}

func FromDomainToConferenceRequest(conference *domain.Conference) *conf.CreateConferenceRequest {
	return &conf.CreateConferenceRequest{
		Name:             conference.Name,
		EventTime:        timestamppb.New(conference.EventTime.Time),
		ParticipantCount: conference.ParticipantCount,
		SpeakerCount:     conference.SpeakerCount,
	}
}

func FromMultiCreateToDomain(list *conf.MultiCreateConferenceRequest) []domain.Conference {
	result := make([]domain.Conference, len(list.Items))
	for i := 0; i < len(list.Items); i++ {
		result[i] = FromRequestToDomain(list.Items[i])
	}
	return result
}

func FromDomainToMultiCreate(source []domain.Conference) *conf.MultiCreateConferenceRequest {
	list := make([]*conf.CreateConferenceRequest, len(source))
	for i := 0; i < len(source); i++ {
		list[i] = FromDomainToConferenceRequest(&source[i])
	}
	return &conf.MultiCreateConferenceRequest{Items: list}
}
