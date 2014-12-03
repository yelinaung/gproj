package main

import (
	"encoding/json"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"net/http"
	"os"
)

type Configuration struct {
	Period   string `json:"period"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	app := cli.NewApp()
	app.Name = "Sitechecker"
	app.Usage = "Site Checker"
	app.Version = "0.1.0"

	app.Action = func(c *cli.Context) {
		args := c.Args()
		if len(args) == 0 {
			println("Error: ", "Argument cannot be empty!")
		} else {
			resp, err := http.Get(args[0])
			if err != nil {
				PanicIf(err)
			} else {
				println("Status ", resp.Status)
			}
		}
	}

	app.Run(os.Args)
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func readConfig() Configuration {
	file, err := ioutil.ReadFile("config.json")
	PanicIf(err)

	config := Configuration{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		PanicIf(err)
	}

	return config
}
