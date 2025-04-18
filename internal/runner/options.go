package runner

import (
	"fmt"

	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
)

// Flag options
type Options struct {
	Input string
	Limit int
	Debug bool
}

func ParseOptions() (*Options, error) {
	options := &Options{}
	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription(`GoTStarter is a scaffold for quickly building terminal tools`)

	flagSet.CreateGroup("input", "Input",
		flagSet.StringVarP(&options.Input, "input", "i", "", ""),
	)

	flagSet.CreateGroup("limit", "Limit",
		flagSet.IntVarP(&options.Limit, "limit", "l", 3, ""),
	)

	flagSet.CreateGroup("debug", "Debug",
		flagSet.BoolVarP(&options.Debug, "debug", "d", false, ""),
	)

	if err := flagSet.Parse(); err != nil {
		return nil, fmt.Errorf("Parse flags error: %v", err)
	}

	// Show banner and version
	ShowBanner()

	// Debug Mode
	if options.Debug {
		gologger.DefaultLogger.SetMaxLevel(levels.LevelDebug)
	}

	// Validate options
	err := options.validateOptions()
	if err != nil {
		return nil, err
	}

	// Display options
	options.outputConfig()

	return options, nil
}

func (o *Options) validateOptions() error {
	//
	if o.Input == "" {
		return fmt.Errorf("Input is required")
	}
	if o.Limit <= 0 {
		return fmt.Errorf("Limit must be greater than 0")
	}
	return nil
}

// Display Key Options
func (o *Options) outputConfig() {
	if o.Debug {
		gologger.Info().Msgf("Debug mode enabled")
	}
}
