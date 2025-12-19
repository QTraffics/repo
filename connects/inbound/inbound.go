package inbound

import (
	"github.com/qtraffics/qtfra/services"
)

type Inbound interface {
	services.LifeCycle

	Type() Type
}
