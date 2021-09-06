package configs

type Config struct {
	Port                     int                      `toml:"port"`
	ConnectionString         string                   `toml:"connection_string"`
	DomainConfiguration      DomainConfiguration      `toml:"domain_configuration"`
	KafkaConfiguration       KafkaConfiguration       `toml:"kafka_configuration"`
	JaegerConfiguration      JaegerConfiguration      `toml:"jaeger_configuration"`
	HealthCheckConfiguration HealthCheckConfiguration `toml:"health_check_configuration"`
	MetricsConfiguration     MetricsConfiguration     `toml:"metrics_configuration"`
}
