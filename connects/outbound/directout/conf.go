package directout

import (
	"fmt"
	"net/netip"
	"strings"

	"github.com/qtraffics/qtfra/ex"
	"github.com/qtraffics/repo/connects/internal/common/stringconf"
	"github.com/qtraffics/repo/connects/outbound"
)

func NewConfFromString(confString string) (*Conf, error) {
	c := new(Conf)
	fields := stringconf.Fields{Raw: confString}
	for k, v := range fields.IterKV() {
		err := c.applyConfiguration(k, v)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

type Conf struct {
	BindInterface string
	BindAddress4  netip.Addr
	BindAddress6  netip.Addr

	ResolverName string
	DNS          string
	EnableTFO    bool
}

func (c *Conf) applyConfiguration(k, v string) error {
	switch strings.ToLower(k) {
	case "tfo":
		c.EnableTFO = true
	case "dns":
		c.DNS = v
	case "resolver":
		c.ResolverName = v
	case "dev":
		c.BindInterface = v
	case "bind4":
		var err error
		c.BindAddress4, err = netip.ParseAddr(v)
		if err != nil {
			return err
		}
	case "bind6":
		var err error
		c.BindAddress6, err = netip.ParseAddr(v)
		if err != nil {
			return err
		}
	default:
		return ex.New("unknown field: ", k, "=", v)
	}
	return nil
}

func (c *Conf) Type() outbound.Type {
	return outbound.TypeDirect
}

func (c *Conf) String() string {
	return fmt.Sprint(*c)
}
