package ova_conference_api

import (
	"context"
	"github.com/rs/zerolog/log"
	conf "ova-conference-api/pkg/api/github.com/ozonva/ova-conference-api/pkg/ova-conference-api"
)

func (server *GRPCServer) CreateConference(ctx context.Context, request *conf.CreateConferenceRequest) (*conf.ConferenceResponse, error) {
	log.Info().Msgf("CreateConference request %v", request)
	result, err := server.repo.AddEntity(ctx, ToConferenceDomain(request))
	if err != nil {
		log.Err(err)
	}
	return ToConferenceResponse(result), err
}
