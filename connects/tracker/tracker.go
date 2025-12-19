package tracker

import (
	"context"
	"net"
)

type Tracker interface {
	NewConnection(ctx context.Context, conn net.Conn) net.Conn
	NewPacketConnection(ctx context.Context, conn net.PacketConn) net.PacketConn
}
