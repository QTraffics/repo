package inbound

import (
	"net"
	"net/netip"

	"github.com/qtraffics/qtfra/services"
)

type Inbound interface {
	services.LifeCycle

	Accept() (net.Conn, error)
	AcceptPacket() ([]byte, netip.AddrPort, error)
}
