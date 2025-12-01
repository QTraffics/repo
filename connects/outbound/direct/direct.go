package direct

import (
	"context"
	"net"
	"os"

	"github.com/qtraffics/qnetwork/addrs"
	"github.com/qtraffics/qnetwork/meta"
	"github.com/qtraffics/repo/connects/outbound"
)

var _ outbound.Outbound = (*Outbound)(nil)

type Outbound struct {
	remote  addrs.Socksaddr
	network meta.Network
}

func (o *Outbound) Connect(ctx context.Context) (net.Conn, error) {
	return nil, os.ErrInvalid
}

func (o *Outbound) ConnectPacket(ctx context.Context) (net.PacketConn, error) {
	return nil, os.ErrInvalid
}

func (o *Outbound) Close() error {
	return nil
}
