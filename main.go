package main

import (
	"fmt"
	_ "github.com/eriknelson/aa-sbx/pkg/broker"
	"github.com/eriknelson/aa-sbx/pkg/config"
	"github.com/eriknelson/aa-sbx/pkg/registry"
	"os"
)

func main() {
	args := config.LoadArgs()
	if args.ConfigFile == "" {
		// TODO args.Usage()
		fmt.Println("Must provide config file --config $FILE")
		os.Exit(1)
	}

	config := config.LoadConfig(args.ConfigFile)

	registry := registry.CreateRegistry(config.Registry)
	registry.LoadApps()
}
