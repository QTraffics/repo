package inbound

import (
	"context"

	"github.com/qtraffics/qtfra/enhancements/iolib"
	"github.com/qtraffics/qtfra/enhancements/pool"
)

var (
	connHandlerContextPool = pool.New[*ConnHandlerContext](func() *ConnHandlerContext {
		return new(ConnHandlerContext)
	})
	packetHandlerContextPool = pool.New[*PacketHandlerContext](func() *PacketHandlerContext {
		return new(PacketHandlerContext)
	})
)

func GetConnContext(ctx context.Context) *ConnHandlerContext {
	get := connHandlerContextPool.Get()
	get.reset()
	get.Context = ctx
	return get
}

func PutConnContext(c *ConnHandlerContext) {
	if c.Conn != nil {
		_ = c.Conn.Close()
	}
	connHandlerContextPool.Put(c)
}

func GetPacketContext() *PacketHandlerContext {
	get := packetHandlerContextPool.Get()
	get.reset()
	return get
}

func PutPacketContext(c *PacketHandlerContext) {
	if c.Writer != nil {
		_ = iolib.Close(c.Writer)
	}
	packetHandlerContextPool.Put(c)
}
