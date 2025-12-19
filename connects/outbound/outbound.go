package outbound

import (
	"context"
	"net"

	"github.com/qtraffics/qnetwork/meta"
	"github.com/qtraffics/qtfra/services"
)

type Outbound interface {
	services.LifeCycle

	Type() Type

	Connect(ctx context.Context, network meta.Network) (net.Conn, error)
}
