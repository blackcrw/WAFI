package main

import (
	"runtime"

	"github.com/blackcrw/akumascan/cli"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	cli.Execute()
}
