package config

import (
	"github.com/projectdiscovery/gologger"
	"github.com/spf13/viper"
)

type Token struct {
	OpenAi   string `yaml:"openai"`
	DeepSeek string `yaml:"deepseek"`
}

type Config struct {
	Token
}

var ConfigInfo Config

// Read config file
func InitConfig() error {
	// Viper config auto find file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		gologger.Error().Msgf("Read config file error：%s", err)
		return err
	}

	if err := viper.Unmarshal(&ConfigInfo); err != nil {
		gologger.Error().Msgf("Parse config file error：%s", err)
		return err
	}

	gologger.Info().Msg("Read Config Successful")
	return nil
}
