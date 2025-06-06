package main

import (
	"github.com/projectdiscovery/gologger"

	"gotstarter/runner"
)

func main() {
	options, err := runner.ParseOptions()
	if err != nil {
		gologger.Error().Msg(err.Error())
		return
	}
	taskRunner, err := runner.New(options)
	if err != nil {
		gologger.Error().Msg(err.Error())
		return
	}
	err = taskRunner.Run()
	if err != nil {
		gologger.Error().Msg(err.Error())
		return
	}
	gologger.Info().Msg("Task Completed")
	return
}
