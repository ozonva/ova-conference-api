package ova_conference_api

import (
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"io"
	"net"
	ova_conference_api "ova-conference-api"
	"ova-conference-api/internal/configs"
	"ova-conference-api/internal/kafka"
	"ova-conference-api/internal/metrics"
	"ova-conference-api/internal/tracer"
	"ova-conference-api/internal/utils/repo"
	conf "ova-conference-api/pkg/api/github.com/ozonva/ova-conference-api/pkg/ova-conference-api"
	"strconv"
)

type ConferenceServer struct {
	conf.ConferencesServer
	repo           repo.Repo
	server         *grpc.Server
	kafkaP         kafka.Producer
	metricsManager metrics.Metrics
	traceCloser    io.Closer
}

func NewServer(repo repo.Repo, service *grpc.Server, producer kafka.Producer, metricsManager metrics.Metrics, traceCloser io.Closer) *ConferenceServer {
	return &ConferenceServer{repo: repo, server: service, kafkaP: producer, metricsManager: metricsManager, traceCloser: traceCloser}
}
func Init(config *configs.Config, metricManager metrics.Metrics) *ConferenceServer {
	log.Log().Msg("Initializing.......")

	repo := repo.NewRepo(config.ConnectionString, config.DomainConfiguration.ChunkSize)
	if err := repo.Open(); err != nil {
		log.Fatal().Msgf("failed to connect to DB : %v", err)
	}

	producer, err := kafka.NewProducer(config.KafkaConfiguration)
	if err != nil {
		log.Fatal().Err(err).Msg("Kafka NewProducer error")
	}

	traceCloser := tracer.InitTracer(ova_conference_api.ApplicationName, &config.JaegerConfiguration)

	service := grpc.NewServer()

	server := NewServer(repo, service, producer, metricManager, traceCloser)
	conf.RegisterConferencesServer(service, server)
	return server
}

func (confServer *ConferenceServer) Start(config *configs.Config, exitChan chan interface{}) {
	go func() {
		log.Log().Msg("Start listening.......")
		listen, err := net.Listen("tcp", ":"+strconv.Itoa(config.Port))
		if err != nil {
			log.Error().Msgf("Failed to listen server %v", err)
			exitChan <- nil
			return
		}
		if err := confServer.server.Serve(listen); err != nil {
			log.Error().Msgf("failed to serve: %v", err)
			exitChan <- nil
		}
	}()
}

func (confServer *ConferenceServer) Stop() {
	var errorList []error
	confServer.server.GracefulStop()

	if err := confServer.repo.Close(); err != nil {
		errorList = append(errorList, err)
		log.Err(err)
	}
	if err := confServer.kafkaP.Close(); err != nil {
		errorList = append(errorList, err)
		log.Err(err)
	}

	if err := confServer.traceCloser.Close(); err != nil {
		errorList = append(errorList, err)
		log.Err(err)
	}

	if len(errorList) > 0 {
		log.Fatal().Msg("Critical error. Server should be fully restarted")
	}
}

func sendNotificationToBroker(server *ConferenceServer, changeType kafka.ChangeType, value interface{}) {
	err := server.kafkaP.Send(struct {
		Type  kafka.ChangeType
		Value interface{}
	}{Type: changeType, Value: value})
	if err != nil {
		log.Err(err).Msg("error while sending message to kafka")
	}
}
