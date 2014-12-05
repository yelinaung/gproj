package main

import (
	"github.com/codegangsta/cli"
	"github.com/yelinaung/gproj/sc/config"
	"net/http"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Sitechecker"
	app.Usage = "Site Checker"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "config.json",
			Usage: "path to config.json",
		},
	}

	app.Action = func(c *cli.Context) {
		if c.String("config") == "config.json" {
			configPath := c.String("config")
			println("Hello", configPath)
			conf := config.ReadConfig(configPath)
			println(conf.Email)

			status := CheckStatus(conf.URL).StatusCode
			if status == http.StatusInternalServerError {
				println("Error!")
			} else {
				println("OK!")
			}
		} else {
			println("Error: ", "Please point to proper config.json")
		}
	}

	app.Run(os.Args)
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckStatus(link string) *http.Response {
	resp, err := http.Get(link)
	PanicIf(err)
	return resp
}
