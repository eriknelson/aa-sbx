package main

import (
	"github.com/jessevdk/go-flags"
)

type Args struct {
	ConfigFile string `short:"c" long:"config" description:"Config File"`
}

func LoadArgs() Args {
	args := Args{}

	_, err := flags.Parse(&args)
	if err != nil {
		panic(err)
	}

	return args
}
