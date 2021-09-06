package configs

type KafkaConfiguration struct {
	Topic   string   `toml:"topic"`
	Brokers []string `toml:"brokers"`
}
