package traffic

import (
	"context"
	"net"
	"sync/atomic"

	"github.com/qtraffics/qtfra/enhancements/iolib/counter"
)

type Metadata struct {
	Upload   *atomic.Int64
	Download *atomic.Int64
}

type Tracker struct {
	Total atomic.Int64
}

func (t *Tracker) NewConnection(ctx context.Context, conn net.Conn) net.Conn {
	upload := new(atomic.Int64)
	download := new(atomic.Int64)

	return &TCPConn{
		Conn: counter.NewCounterConn(conn, []counter.Func{func(n int64) {
			download.Add(n)
			t.Total.Add(n)
		}}, []counter.Func{
			func(n int64) {
				upload.Add(n)
				t.Total.Add(n)
			},
		}),
		metadata: Metadata{
			Upload:   upload,
			Download: download,
		},
	}
}

func (t *Tracker) NewPacketConnection(ctx context.Context, conn net.PacketConn) net.PacketConn {
	upload := new(atomic.Int64)
	download := new(atomic.Int64)

	return &UDPConn{
		PacketConn: counter.NewCounterPacketConn(conn, []counter.Func{func(n int64) {
			download.Add(n)
			t.Total.Add(n)
		}}, []counter.Func{
			func(n int64) {
				upload.Add(n)
				t.Total.Add(n)
			},
		}),
		metadata: Metadata{
			Upload:   upload,
			Download: download,
		},
	}
}

func GetMetadata(c any) (Metadata, bool) {
	type trafficMetadata interface {
		TrafficMetadata() Metadata
	}

	if tc, ok := c.(trafficMetadata); ok {
		return tc.TrafficMetadata(), true
	}
	return Metadata{}, false
}
