package ova_conference_api

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"ova-conference-api/internal/utils"
	conf "ova-conference-api/pkg/api/github.com/ozonva/ova-conference-api/pkg/ova-conference-api"
	"strconv"
)

type GRPCServer struct {
	conf.ConferencesServer
}

func NewServer() *GRPCServer {
	return &GRPCServer{}
}

func Start(config utils.Config) error {
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(config.Port))
	if err != nil {
		log.Fatalf("Failed to listen server %v", err)
	}

	service := grpc.NewServer()
	conf.RegisterConferencesServer(service, NewServer())

	if err := service.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return nil
}
