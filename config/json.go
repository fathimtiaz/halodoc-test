package config

import (
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
)

func getJson() (cfg Config, err error) {

	var (
		data []byte
	)

	if data, err = ioutil.ReadFile("config.json"); err != nil {
		return Config{}, err
	}

	if err = jsoniter.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
