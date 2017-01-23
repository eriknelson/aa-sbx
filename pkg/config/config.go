package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type RegistryConfig struct {
	Name string
	Url  string
}

type Config struct {
	Registry   RegistryConfig
	ConfigFile string
}

func LoadConfig(configFile string) *Config {
	var err error

	config := Config{
		ConfigFile: configFile,
	}

	// TODO: Confirm it exists
	// TODO: Run validator

	// Load struct
	dat, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(dat, &config)
	if err != nil {
		panic(err)
	}

	return &config
}
