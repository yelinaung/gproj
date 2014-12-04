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

	app.Action = func(c *cli.Context) {
		args := c.Args()
		if len(args) == 0 {
			println("Error: ", "Argument cannot be empty!")
		} else {
			status := CheckStatus(args[0]).StatusCode
			if status == http.StatusInternalServerError {
				println("Error!")
			} else {
				println("OK!")
			}
		}
	}

	c := config.ReadConfig()
	println(c.Email)
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
