package configs

type JaegerConfiguration struct {
	Host string `toml:"Host"`
	Port int    `toml:"port"`
}
