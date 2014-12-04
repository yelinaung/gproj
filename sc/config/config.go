package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Period   string `json:"period"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ReadConfig() Config {
	file, err := ioutil.ReadFile("config.json")
	PanicIf(err)

	config := Config{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		PanicIf(err)
	}

	return config
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}
