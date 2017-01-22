package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Registry struct {
		Type string
		Url  string
	}
	ConfigFile string
}

func LoadConfig(configFile string) Config {
	var err error

	config := Config{
		ConfigFile: configFile,
	}

	dat, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(dat, &config)
	if err != nil {
		panic(err)
	}

	return config
}
