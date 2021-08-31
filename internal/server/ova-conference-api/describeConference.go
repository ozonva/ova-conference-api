package ova_conference_api

import (
	"context"
	"github.com/rs/zerolog/log"
	conf "ova-conference-api/pkg/api/github.com/ozonva/ova-conference-api/pkg/ova-conference-api"
)

func (server *GRPCServer) DescribeConference(ctx context.Context, request *conf.EntityConferenceRequest) (*conf.ConferenceResponse, error) {
	log.Info().Msgf("DescribeConference request: %v", request)
	result, err := server.repo.DescribeEntity(ctx, request.Id)
	if err != nil {
		log.Err(err)
		return nil, err
	}
	return ToConferenceResponse(result), err
}
