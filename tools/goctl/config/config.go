package config

import (
	_ "embed"
)

const (
	DefaultFormat = "gozero"
	configFile    = "goctl.yaml"
)

//go:embed default.yaml
var defaultConfig []byte

type (
	Config struct {
		NamingFormat string `yaml:"namingFormat"`
	}

	External struct {
		Model Model `yaml:"model,omitempty"`
	}

	Model struct {
		TypesMap map[string]ModelTypeMapOption `yaml:"types_map,omitempty" `
	}

	ModelTypeMapOption struct {
		Type string `yaml:"type"`

		UnsignedType string `yaml:"unsigned_type,omitempty"`

		NullType string `yaml:"null_type,omitempty"`

		Pkg string `yaml:"pkg,omitempty"`
	}
)

func NewConfig(format string) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

func GetExternalConfig() (*External, error) { _ = "STUB: not implemented"; return nil, nil }

func loadConfig(cfg *External) error { _ = "STUB: not implemented"; return nil }

func getConfigPath(workDir string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func validate(cfg *Config) error { _ = "STUB: not implemented"; return nil }
