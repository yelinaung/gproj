package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.google.com")
	PanicIf(err)
	fmt.Println(resp.Status)
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}
