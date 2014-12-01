package main

import (
	"github.com/codegangsta/cli"
	"net/http"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "sitechecker"
	app.Usage = "Site Checker"

	app.Action = func(c *cli.Context) {
		resp, err := http.Get(c.Args()[0])
		PanicIf(err)
		println("Status ", resp.Status)
	}

	app.Run(os.Args)
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}
