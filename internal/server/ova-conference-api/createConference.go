package ova_conference_api

import (
	"context"
	"github.com/rs/zerolog/log"
	"ova-conference-api/internal/kafka"
	conf "ova-conference-api/pkg/api/github.com/ozonva/ova-conference-api/pkg/ova-conference-api"
)

func (confServer *ConferenceServer) CreateConference(ctx context.Context, request *conf.CreateConferenceRequest) (*conf.ConferenceEntity, error) {
	log.Info().Msgf("CreateConference request: %v", request)
	result, err := confServer.repo.AddEntity(ctx, FromRequestToDomain(request))
	if err != nil {
		log.Err(err)
		return nil, err
	}
	confServer.metricsManager.CreateConferenceEvent()
	sendNotificationToBroker(confServer, kafka.CreateConference, *result)

	return FromDomainToConferenceEntity(result), err
}
