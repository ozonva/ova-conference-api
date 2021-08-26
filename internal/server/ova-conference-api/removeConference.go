package ova_conference_api

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	conf "ova-conference-api/pkg/api/github.com/ozonva/ova-conference-api/pkg/ova-conference-api"
)

func (G *GRPCServer) RemoveConference(ctx context.Context, request *conf.CreateConferenceRequest) (*empty.Empty, error) {
	panic("implement me")
}
