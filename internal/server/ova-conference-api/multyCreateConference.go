package ova_conference_api

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"ova-conference-api/internal/domain"
	"ova-conference-api/internal/kafka"
	conf "ova-conference-api/pkg/api/github.com/ozonva/ova-conference-api/pkg/ova-conference-api"
)

func (confServer *ConferenceServer) MultiCreateConference(ctx context.Context, request *conf.MultiCreateConferenceRequest) (*conf.ListConferenceEntity, error) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("MultiCreateConference")
	defer span.Finish()

	bulkCallback := buildBulkCallBack(tracer, span)

	result, err := confServer.repo.AddEntities(ctx, FromMultiCreateToDomain(request), bulkCallback)
	if err != nil {
		log.Err(err)
	}
	response := &conf.ListConferenceEntity{}
	mapItems(response, result)
	confServer.metricsManager.MultiCreateConferenceEvent()
	sendNotificationToBroker(confServer, kafka.MultiCreateConference, result)
	return response, err
}

func buildBulkCallBack(tracer opentracing.Tracer, span opentracing.Span) func(chunk []domain.Conference) {
	return func(chunk []domain.Conference) {
		childSpan := tracer.StartSpan(
			"MultiCreateConference: chunk",
			opentracing.Tag{Key: "chunkSize", Value: len(chunk)},
			opentracing.ChildOf(span.Context()),
		)
		childSpan.Finish()
	}
}
