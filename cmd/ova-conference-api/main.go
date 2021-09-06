package main

import (
	"context"
	"flag"
	"github.com/peak/go-config"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"ova-conference-api/internal/configs"
	"ova-conference-api/internal/healthCheck"
	"ova-conference-api/internal/logger"
	"ova-conference-api/internal/metrics"
	server "ova-conference-api/internal/server/ova-conference-api"
	"syscall"
)

var (
	ConferenceServer  *server.ConferenceServer
	MetricManager     metrics.Metrics
	HealthCheckServer *healthCheck.Server
	MetricsServer     *metrics.Server
)

func main() {
	logger.SetUpLogger()
	log.Print("Configuring server...")
	ctx, cancelFunc := context.WithCancel(context.Background())

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
		<-ch
		log.Warn().Msg("Got signal, cleaning up...")
		cancelFunc()
		stopApplication()
	}()

	filename := flag.String("config", "test_config.toml", "Config filename")

	config, err := configs.ReadConfigFromFile(filename)()
	if err != nil {
		log.Fatal().Err(err)
	}
	HealthCheckServer = healthCheck.NewHealthCheckServer(config)
	MetricsServer = metrics.NewMetricsServer(config)
	MetricManager = metrics.NewMetrics()

	log.Printf("Config is %v", config)
	exitChan := make(chan interface{}, 1)

	go watch(ctx, *filename, exitChan)
	startApplication(config, exitChan)

	for {
		<-exitChan
		log.Log().Msg("Exiting....")
		return
	}

}

func watch(ctx context.Context, filename string, exitChan chan interface{}) {
	configChan, err := config.Watch(ctx, filename)
	if err != nil {
		panic(err)
	}
	for {
		select {
		case e := <-configChan:
			if e != nil {
				log.Error().Msgf("Error occured watching file: %v", e)
				continue
			}
			log.Info().Msg("Changed, reloading...")
			var cfg configs.Config
			err := config.Load(filename, &cfg)
			log.Info().Msgf("Loaded: %v %#v\n", err, cfg)
			stopApplication()
			log.Info().Msg("Server successfully stopped\n")
			startApplication(&cfg, exitChan)
			log.Info().Msg("Server successfully reloaded\n")
		case <-ctx.Done():
			return
		}
	}
}

func startApplication(config *configs.Config, exitChan chan interface{}) {
	log.Print("Starting server...")
	service := server.Init(config, MetricManager)
	HealthCheckServer.Start()
	MetricsServer.Start()
	ConferenceServer = service
	service.Start(config, exitChan)
}

func stopApplication() {
	HealthCheckServer.Stop()
	MetricsServer.Stop()
	ConferenceServer.Stop()
}
