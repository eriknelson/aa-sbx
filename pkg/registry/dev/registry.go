package devreg

import (
	"fmt"
	"github.com/eriknelson/aa-sbx/pkg/config"
)

type DevRegistry struct {
	config config.RegistryConfig
}

func (r *DevRegistry) Init(config config.RegistryConfig) error {
	fmt.Printf("DevRegistry::Init with url -> [ %s ] \n", config.Url)
	r.config = config
	return nil
}

func (r *DevRegistry) LoadApps() error {
	fmt.Println("DevRegistry::LoadApps ")
	return nil
}
