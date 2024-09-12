package main

import (
	"github.com/xnacly/libjson"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		os.Exit(-1)
	}

	_, err = libjson.NewReader(f)
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
