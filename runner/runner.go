package runner

import (
	"github.com/projectdiscovery/gologger"

	"gotstarter/internal/initialize"
)

type Runner struct {
	Options *Options
}

func New(options *Options) (*Runner, error) {
	runner := &Runner{
		Options: options,
	}
	err := initialize.Init()
	if err != nil {
		return nil, err
	}
	return runner, nil
}

func (r *Runner) Run() error {
	gologger.Info().Msg(r.Options.Input)
	return nil
}
