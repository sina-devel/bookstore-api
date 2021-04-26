package config

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
	"os"
)

var (
	cfg *Config
)

func ReadFile(cfg *Config) (err error) {
	cfgPath := "./build/config/config.yaml"

	file, err := os.Open(cfgPath)
	if err != nil {
		return err
	}

	defer func() {
		cerr := file.Close()
		if err != nil {
			err = cerr
		}
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return err
	}

	return nil
}

func ReadEnv(cfg *Config) error {
	return envconfig.Process("", cfg)
}

func SetConfig(c *Config) {
	cfg = c
}
