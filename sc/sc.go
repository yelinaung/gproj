package main

import (
	"github.com/codegangsta/cli"
	"github.com/yelinaung/gproj/sc/config"
	"net/http"
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
			resp, err := http.Get(args[0])
			if err != nil {
				PanicIf(err)
			} else {
				println("Status ", resp.Status)
			}
		}
	}

	// app.Run(os.Args)
	//config := config.Config{}
	c := config.ReadConfig()
	println(c.Email)
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckStatus(link string) http.Status {
	resp, err := http.Get(link)
}
