package main

import (
	"fmt"
	_ "github.com/eriknelson/aa-sbx/pkg/broker"
	_ "github.com/eriknelson/aa-sbx/pkg/registry"
)

func main() {
	args := LoadArgs()
	config := LoadConfig(args.ConfigFile)
	fmt.Printf("Reading ConfigFile: %s\n", args.ConfigFile)
	fmt.Printf("Registry type: %s\n", config.Registry.Type)
	fmt.Printf("Registry url: %s\n", config.Registry.Url)
}
