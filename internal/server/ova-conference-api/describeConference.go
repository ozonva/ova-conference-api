package ova_conference_api

import (
	"context"
	"github.com/rs/zerolog/log"
	conf "ova-conference-api/pkg/api/github.com/ozonva/ova-conference-api/pkg/ova-conference-api"
)

func (G *GRPCServer) DescribeConference(ctx context.Context, request *conf.CreateConferenceRequest) (*conf.ConferenceResponse, error) {
	log.Info().Msg("DescribeConference request")
	return getDummyConference(), nil
}
