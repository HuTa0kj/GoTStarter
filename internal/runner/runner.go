package runner

import (
	"github.com/projectdiscovery/gologger"
	"gotstarter/internal/initialize"
)

func Runner() error {
	options, err := ParseOptions()
	if err != nil {
		return err
	}
	err = initialize.Init()
	if err != nil {
		return err
	}
	gologger.Info().Msg(options.Input)
	return nil
}
