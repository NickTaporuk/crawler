package main

import (
	"github.com/NickTaporuk/crawler/parser"
	"github.com/NickTaporuk/crawler/runner"
	"runtime"
	"time"
)


const (
	// timeout use to cancel process if runner exceeded time limit
	timeout = 2 * time.Minute
)

func main() {
	//fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	//
	r := runner.New(timeout)

	//r.Add(parser.LunParser())
	r.Add(
		parser.LunParser(),
		parser.GoogleSearchParser(),
		)

	err := r.Start()
	if err != nil {
		panic(err)
	}
}
