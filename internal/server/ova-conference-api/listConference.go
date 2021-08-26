package ova_conference_api

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rs/zerolog/log"
	conf "ova-conference-api/pkg/api/github.com/ozonva/ova-conference-api/pkg/ova-conference-api"
)

func (G *GRPCServer) ListConference(ctx context.Context, req *empty.Empty) (*conf.ListConferenceResponse, error) {
	log.Info().Msg("ListConference request")
	return getDummyList(), nil
}

func getDummyList() *conf.ListConferenceResponse {
	items := []*conf.ConferenceResponse{getDummyConference()}
	return &conf.ListConferenceResponse{Items: items}
}
