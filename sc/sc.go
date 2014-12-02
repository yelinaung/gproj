package main

import (
	"github.com/codegangsta/cli"
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
			resp, err := http.Get(args[0])
			if err != nil {
				PanicIf(err)
			}
			println("Status ", resp.Status)
		}
	}

	app.Run(os.Args)
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}
