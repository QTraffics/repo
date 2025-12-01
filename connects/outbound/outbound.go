package outbound

import (
	"context"
	"net"
)

type Outbound interface {
	Connect(ctx context.Context) (net.Conn, error)
	ConnectPacket(ctx context.Context) (net.PacketConn, error)

	Close() error
}
