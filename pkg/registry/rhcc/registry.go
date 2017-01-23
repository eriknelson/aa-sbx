package rhccreg

import (
	"fmt"
	"github.com/eriknelson/aa-sbx/pkg/config"
)

type RHCCRegistry struct {
	config config.RegistryConfig
}

func (r *RHCCRegistry) Init(config config.RegistryConfig) error {
	fmt.Printf("RHCCRegistry::Init with url -> [ %s ] \n", config.Url)
	r.config = config
	return nil
}

func (r *RHCCRegistry) LoadApps() error {
	fmt.Println("RHCCRegistry::LoadApps ")
	return nil
}
