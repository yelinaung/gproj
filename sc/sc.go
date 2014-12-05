package main

import (
	"github.com/codegangsta/cli"
	"github.com/rakyll/ticktock"
	"github.com/rakyll/ticktock/t"
	"github.com/yelinaung/gproj/sc/config"
	"net/http"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Sitechecker"
	app.Usage = "Just a configurable Site Checker"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "config.json",
			Usage: "path to config.json",
		},
	}

	app.Action = func(c *cli.Context) {
		if c.String("config") == "config.json" || c.String("c") == "config.json" {
			configPath := c.String("config")
			conf := config.ReadConfig(configPath)

			println("URL is : ", conf.URL)

			err := ticktock.Schedule(
				"site check",
				&CheckJob{
					conf.URL,
				},
				&t.When{
					Every: t.Every(conf.Interval).Seconds()})
			ticktock.Start()
			PanicIf(err)
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

type CheckJob struct {
	URL string
}

func (cj *CheckJob) Run() error {
	status := CheckStatus(cj.URL).StatusCode
	if status == http.StatusInternalServerError {
		println("Status : ", "Error!")
	} else {
		println("Status : ", "OK!")
	}
	return nil
}
