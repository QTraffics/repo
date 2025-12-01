package direct

import (
	"github.com/qtraffics/qnetwork/addrs"
	"github.com/qtraffics/repo/connects/outbound"
)

func NewConfString(confString string) (Conf, error) {
	return Conf{}, nil
}

var _ outbound.Conf = (*Conf)(nil)

type Conf struct {
	TFO bool
	DNS addrs.Socksaddr
}

func (c *Conf) Apply(outbound outbound.Outbound) error {
	directOutbound := outbound.(*Outbound)
	_ = directOutbound
	return nil
}

func (c *Conf) String() string {
	return ""
}
