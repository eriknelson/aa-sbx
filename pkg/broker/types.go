package broker

import (
	"github.com/pborman/uuid"
)

type AppSpec struct {
	Id       string
	Name     string
	Bindable bool
	Async    string // required, optional, unsupported
}
