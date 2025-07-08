package runner

import (
	"github.com/projectdiscovery/gologger"
)

type Runner struct {
	Options *Options
}

func New(options *Options) (*Runner, error) {
	runner := &Runner{
		Options: options,
	}
	return runner, nil
}

func (r *Runner) Run() error {
	gologger.Info().Msg(r.Options.Input)
	return nil
}
