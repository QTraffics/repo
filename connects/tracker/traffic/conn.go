package traffic

import (
	"net"
)

type TCPConn struct {
	net.Conn

	metadata Metadata
}

func (c *TCPConn) TrafficMetadata() Metadata {
	return c.metadata
}

type UDPConn struct {
	net.PacketConn

	metadata Metadata
}

func (c *UDPConn) TrafficMetadata() Metadata {
	return c.metadata
}
