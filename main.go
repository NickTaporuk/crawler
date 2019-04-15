package main

import (
	"github.com/NickTaporuk/crawler/parser"
	"github.com/NickTaporuk/crawler/runner"
	"time"
)


const (
	// timeout use to cancel process if runner exceeded time limit
	timeout = 2 * time.Minute
)

func main() {
	//
	r := runner.New(timeout)

	r.Add(parser.LunParser())

	err := r.Start()
	if err != nil {
		panic(err)
	}
}
