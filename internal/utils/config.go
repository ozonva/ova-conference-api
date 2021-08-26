package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Port             int `json:"port"`
	TimeoutInSeconds int `json:"timeout_in_seconds"`
}

func ReadConfigFromFile(filePah string) func() (Config, error) {
	openConfiguration := func() (Config, error) {
		cfg := Config{}
		file, err := os.Open(filePah)
		if err != nil {
			return cfg, err
		}
		defer file.Close()
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			return cfg, err
		}
		err = json.Unmarshal(bytes, &cfg)
		return cfg, err
	}
	return openConfiguration
}

func ReadConfigInLoop(filePath string, loopCnt int) []Config {
	openConfigFunc := ReadConfigFromFile(filePath)
	result := make([]Config, loopCnt)
	for i := 0; i < loopCnt; i++ {
		cfg, err := openConfigFunc()
		if err != nil {
			fmt.Println(err)
			continue
		}
		result[i] = cfg
	}
	return result
}
