package directout

import (
	"context"
	"net"
	"os"

	"github.com/qtraffics/qnetwork/addrs"
	"github.com/qtraffics/qnetwork/meta"
)

type Outbound struct {
	Conf

	remote addrs.Socksaddr
}

func (o *Outbound) Start(ctx context.Context) error {
	return nil
}

func (o *Outbound) Connect(ctx context.Context, network meta.Network) (net.Conn, error) {
	return nil, os.ErrInvalid
}

func (o *Outbound) Close() error {
	return nil
}
