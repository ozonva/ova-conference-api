package ova_conference_api

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rs/zerolog/log"
	"ova-conference-api/internal/kafka"
	conf "ova-conference-api/pkg/api/github.com/ozonva/ova-conference-api/pkg/ova-conference-api"
)

func (confServer *ConferenceServer) UpdateConference(ctx context.Context, request *conf.ConferenceEntity) (*empty.Empty, error) {
	log.Info().Msgf("DescribeConference request: %v", request)
	domainEntity := FromConferenceEntityToDomain(request)
	err := confServer.repo.UpdateEntity(ctx, domainEntity)
	if err != nil {
		log.Err(err)
		return nil, err
	}

	confServer.metricsManager.CreateConferenceEvent()
	sendNotificationToBroker(confServer, kafka.UpdateConference, domainEntity)
	return nil, err
}
