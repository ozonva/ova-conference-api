package ova_conference_api

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rs/zerolog/log"
	conf "ova-conference-api/pkg/api/github.com/ozonva/ova-conference-api/pkg/ova-conference-api"
)

func (server *GRPCServer) RemoveConference(ctx context.Context, request *conf.EntityConferenceRequest) (*empty.Empty, error) {
	log.Info().Msgf("RemoveConference request %v", request)
	err := server.repo.DeleteEntity(ctx, request.Id)
	if err != nil {
		log.Err(err)
	}

	return nil, err
}
