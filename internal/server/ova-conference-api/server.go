package ova_conference_api

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"ova-conference-api/internal/utils"
	"ova-conference-api/internal/utils/repo"
	conf "ova-conference-api/pkg/api/github.com/ozonva/ova-conference-api/pkg/ova-conference-api"
	"strconv"
)

type GRPCServer struct {
	conf.ConferencesServer
	repo repo.Repo
}

func NewServer(repo repo.Repo) *GRPCServer {
	return &GRPCServer{repo: repo}
}

func Start(config utils.Config) error {
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(config.Port))
	if err != nil {
		log.Fatalf("Failed to listen server %v", err)
	}

	repo := repo.NewRepo(config.ConnectionString)
	if err := repo.Open(); err != nil {
		log.Fatalf("failed to connect to DB : %v", err)
	}

	service := grpc.NewServer()
	conf.RegisterConferencesServer(service, NewServer(repo))

	if err := service.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return nil
}
