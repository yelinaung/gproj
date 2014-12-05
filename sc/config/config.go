package config

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Config struct {
	Period   time.Duration `json:"period"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
	URL      string        `json:"url"`
}

func ReadConfig(filename string) Config {
	file, err := ioutil.ReadFile(filename)
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
