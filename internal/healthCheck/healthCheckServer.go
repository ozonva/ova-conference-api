package healthCheck

import (
	"context"
	"github.com/etherlabsio/healthcheck/v2"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"net/http"
	"ova-conference-api/internal/configs"
	"strconv"
	"sync"
	"time"
)

type Server struct {
	httpServer *http.Server
	wg         *sync.WaitGroup
	config     *configs.Config
}

func NewHealthCheckServer(configuration *configs.Config) *Server {
	return &Server{
		config: configuration,
	}
}

func (s *Server) Start() {
	mux := http.NewServeMux()
	db, _ := sqlx.Open("postgres", s.config.ConnectionString)
	mux.Handle("/healthcheck", healthcheck.Handler(
		healthcheck.WithTimeout(5*time.Second),

		healthcheck.WithChecker(
			"database", healthcheck.CheckerFunc(
				func(ctx context.Context) error {
					return db.PingContext(ctx)
				},
			),
		),
	))

	s.httpServer = &http.Server{
		Addr:    ":" + strconv.Itoa(s.config.HealthCheckConfiguration.Port),
		Handler: mux,
	}

	s.wg = &sync.WaitGroup{}
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		if err := s.httpServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Err(err).Msg("HealthCheck server: failed to serve")
		}

	}()
}

func (s *Server) Stop() {
	if err := s.httpServer.Shutdown(context.Background()); err != nil {
		log.Err(err).Msg("HealthCheck server: failed to shutdown")
	}
	s.wg.Wait()
}
