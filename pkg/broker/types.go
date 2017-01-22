package broker

import (
	"github.com/pborman/uuid"
)

type AsyncSupport int

const (
	AsyncRequired AsyncSupport = iota
	AsyncOptional
	AsyncUnsupported
)

type AppMeta struct {
	Id       uuid.UUID
	Name     string
	Bindable bool
	Async    AsyncSupport
}
