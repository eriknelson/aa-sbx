package registry

type Registry interface {
	Init( /*url*/ string)
	LoadApps() error
}
