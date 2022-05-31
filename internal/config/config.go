package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

// Config represents the config file written in yaml format.
type Config struct {
	RPC string `yaml:"rpc"`
}

// Load loads the config from file.
// If path is empty, it will search for the config file in the following order:
//
// - /etc/derocli/derocli.yaml
//
// - $HOME/.config/derocli/derocli.yaml
//
// - derocli.yaml
//
// If the config file is not found, it will return an empty config and nil error.
func Load(path string) (cfg *Config, err error) {
	if path != "" {
		return load(path)
	}
	for _, f := range [4]string{
		".derocli",
		"derocli",
	} {
		cfg, err = load(f)
		if err != nil && os.IsNotExist(err) {
			err = nil
			continue
		} else if err != nil && errors.As(err, &viper.ConfigFileNotFoundError{}) {
			err = nil
			continue
		}
	}
	return
}

func load(file string) (cfg *Config, err error) {
	viper.SetConfigName(file)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/derocli/")
	viper.AddConfigPath("$HOME/.config/derocli")
	viper.AddConfigPath(".")
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	if err = viper.Unmarshal(&cfg); err != nil {
		return
	}
	return
}
