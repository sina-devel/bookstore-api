package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

var (
	cfg *Config
	// ErrUnknownFileExtension is returned by the Parse function
	// when the file extension is not allowed for configuration
	ErrUnknownFileExtension = errors.New("unknown file extension")
)

// Parse parses config file into Config
func Parse(path string, cfg *Config) error {
	switch filepath.Ext(path) {
	case ".yaml", ".yml":
		return parseYAML(path, cfg)
	default:
		return ErrUnknownFileExtension
	}
}

// ReadEnv reads some configs from environment variables
func ReadEnv(cfg *Config) error {
	return envconfig.Process("", cfg)
}

// SetConfig sets cfg in config package
func SetConfig(c *Config) {
	cfg = c
}

// parseYAML parses yaml config file into Config
func parseYAML(path string, cfg *Config) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer func() {
		if e := file.Close(); err == nil {
			err = e
		}
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return err
	}

	return nil
}
