package registry

import (
	"github.com/eriknelson/aa-sbx/pkg/config"
	dev "github.com/eriknelson/aa-sbx/pkg/registry/dev"
	rhcc "github.com/eriknelson/aa-sbx/pkg/registry/rhcc"
)

type Registry interface {
	Init(config.RegistryConfig) error
	LoadApps() error
}

func CreateRegistry(config config.RegistryConfig) Registry {
	var reg Registry

	switch config.Name {
	case "dev":
		reg = &dev.DevRegistry{}
	case "rhcc":
		reg = &rhcc.RHCCRegistry{}
	default:
		panic("Unknown registry")
	}

	reg.Init(config)

	return reg
}
