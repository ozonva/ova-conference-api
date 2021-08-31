package ova_conference_api

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"ova-conference-api/internal/domain"
	conf "ova-conference-api/pkg/api/github.com/ozonva/ova-conference-api/pkg/ova-conference-api"
)

func (server *GRPCServer) ListConference(ctx context.Context, request *conf.ListConferenceRequest) (*conf.ListConferenceResponse, error) {
	log.Info().Msgf("ListConference request %v", request)
	if request.Limit < 0 || request.Offset < 0 {
		log.Log().Msg("Request param are incorrect")
		return nil, errors.New("offset and limit should be zero or greater")
	}
	result, err := server.repo.ListEntities(ctx, request.Limit, request.Offset)
	if err != nil {
		log.Err(err)
	}
	response := &conf.ListConferenceResponse{}
	mapItems(response, result)
	return response, err
}

func mapItems(response *conf.ListConferenceResponse, elements []domain.Conference) {
	result := make([]*conf.ConferenceResponse, len(elements))
	for idx, val := range elements {
		result[idx] = ToConferenceResponse(&val)
	}
	response.Items = result
}
