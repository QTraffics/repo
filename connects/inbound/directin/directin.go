package directin

import (
	"context"
	"net"
	"net/netip"

	"github.com/qtraffics/repo/connects/inbound"
)

var _ inbound.Inbound = (*Inbound)(nil)

type Inbound struct{}

func (i *Inbound) Start(ctx context.Context) error {
	// TODO implement me
	panic("implement me")
}

func (i *Inbound) Close() error {
	// TODO implement me
	panic("implement me")
}

func (i *Inbound) Accept() (net.Conn, error) {
	// TODO implement me
	panic("implement me")
}

func (i *Inbound) AcceptPacket() ([]byte, netip.AddrPort, error) {
	// TODO implement me
	panic("implement me")
}
