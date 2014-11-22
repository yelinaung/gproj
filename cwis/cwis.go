package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func wordCountFromFile(fname string) []string {
	str, err := ioutil.ReadFile(fname)
	PanicIf(err)
	return strings.Fields(string(str))
}

func main() {
	arg := os.Args[1]
	total := len(wordCountFromFile(arg))
	fmt.Printf("total words : %d\n", total)
}
