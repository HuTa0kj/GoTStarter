package main

import (
	"github.com/projectdiscovery/gologger"

	"gotstarter/internal/runner"
)

func main() {
	err := runner.Runner()
	if err != nil {
		gologger.Error().Msg(err.Error())
	}
}
