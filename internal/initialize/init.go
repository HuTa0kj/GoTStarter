package initialize

import (
	"github.com/projectdiscovery/gologger"
	"gotstarter/config"
)

// Init All
func Init() error {
	err := config.InitConfig()
	if err != nil {
		return err
	}
	gologger.Debug().Msg("Initialization completed")
	return nil
}
