package metrics

import (
	"context"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"net/http"
	"ova-conference-api/internal/configs"
	"strconv"
	"sync"
)

type Server struct {
	config     *configs.Config
	httpServer *http.Server
	wg         *sync.WaitGroup
}

func NewMetricsServer(configuration *configs.Config) *Server {
	return &Server{
		config: configuration,
	}
}

func (s *Server) Start() {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	s.httpServer = &http.Server{
		Addr:    ":" + strconv.Itoa(s.config.MetricsConfiguration.Port),
		Handler: mux,
	}

	s.wg = &sync.WaitGroup{}
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		if err := s.httpServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Err(err).Msg("Metrics server: failed to serve")
		}
	}()
}

func (s *Server) Stop() {
	if err := s.httpServer.Shutdown(context.Background()); err != nil {
		log.Err(err).Msg("Metrics server: failed to shutdown")
	}
	s.wg.Wait()
}
