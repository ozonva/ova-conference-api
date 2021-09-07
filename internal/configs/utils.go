package configs

import (
	"flag"
	"fmt"
	"github.com/peak/go-config"
)

func ReadConfigFromFile(filename *string) func() (*Config, error) {
	openConfiguration := func() (*Config, error) {
		var cfg Config
		err := config.Load(*filename, &cfg)
		return &cfg, err
	}
	return openConfiguration
}

func ReadConfigInLoop(filePath string, loopCnt int) []Config {
	fileName := flag.String("config", filePath, "Config filename")
	openConfigFunc := ReadConfigFromFile(fileName)
	result := make([]Config, loopCnt)
	for i := 0; i < loopCnt; i++ {
		cfg, err := openConfigFunc()
		if err != nil {
			fmt.Println(err)
			continue
		}
		result[i] = *cfg
	}
	return result
}
