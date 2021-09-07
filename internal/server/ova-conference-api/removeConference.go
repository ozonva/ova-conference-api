package ova_conference_api

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rs/zerolog/log"
	"ova-conference-api/internal/kafka"
	conf "ova-conference-api/pkg/api/github.com/ozonva/ova-conference-api/pkg/ova-conference-api"
)

func (confServer *ConferenceServer) RemoveConference(ctx context.Context, request *conf.EntityConferenceRequest) (*empty.Empty, error) {
	log.Info().Msgf("RemoveConference request: %v", request)
	err := confServer.repo.DeleteEntity(ctx, request.Id)
	if err != nil {
		log.Err(err)
	}
	confServer.metricsManager.DeleteConferenceEvent()
	sendNotificationToBroker(confServer, kafka.DeleteConference, request.Id)
	return nil, err
}
