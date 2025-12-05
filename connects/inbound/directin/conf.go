package directin

import (
	"fmt"
	"net"
	"net/netip"
	"strconv"
	"strings"

	"github.com/qtraffics/qnetwork/addrs"
	"github.com/qtraffics/qnetwork/meta"
	"github.com/qtraffics/qtfra/ex"
	"github.com/qtraffics/repo/connects/inbound"
	"github.com/qtraffics/repo/connects/internal/common/stringconf"
)

type Conf struct {
	Network meta.ProtocolList
	Port    uint16

	BindInterface string
	BindAddress   string
	BindAddress4  netip.Addr
	BindAddress6  netip.Addr
	EnableTFO     bool
	EnableUDPFrag bool

	ReuseAddr bool
	ReusePort bool
}

func NewConfFromString(confString string) (*Conf, error) {
	c := new(Conf)
	if idx := strings.Index(confString, "://"); idx > 0 {
		protocol := meta.ParseProtocolList(confString[:idx])
		if len(protocol) == 0 {
			return nil, ex.New("unknown protocol:", confString[:idx])
		}
		c.Network = protocol
		// trim
		confString = confString[idx+len("://"):]
	} else if idx == 0 {
		return nil, ex.New("unexcepted behavior: start with '://'")
	} else {
		c.Network = meta.ProtocolList{meta.ProtocolTCP, meta.ProtocolUDP}
	}

	fields := stringconf.Fields{Raw: confString}
	// listen
	listen := fields.PickNext()
	if strings.Contains(listen, ":") {
		host, port, err := net.SplitHostPort(listen)
		if err != nil {
			return nil, ex.Cause(err, "SplitHostPort")
		}

		if err = ex.Errors(c.applyHost(host), c.applyPort(port)); err != nil {
			return nil, err
		}
	} else if port, err := strconv.ParseUint(listen, 10, 16); err == nil {
		c.Port = uint16(port)
	} else {
		return nil, ex.New("wrong listen config: ", listen)
	}

	for k, v := range fields.IterKV() {
		if err := c.applyExtra(k, v); err != nil {
			return nil, ex.Cause(err, "extra")
		}
	}
	return c, nil
}

func (c *Conf) applyExtra(k, v string) error {
	switch strings.ToLower(k) {
	case "tfo":
		c.EnableTFO = true
	case "udp_fragment":
		c.EnableUDPFrag = true
	case "reuseaddr":
		c.ReuseAddr = true
	default:
		return ex.New("unknown configuration: ", k+"="+v)
	}
	return nil
}

func (c *Conf) applyHost(host string) error {
	if host == "" {
		c.BindAddress = ""
	} else if _, err := net.InterfaceByName(host); err == nil {
		c.BindInterface = host
	} else if addr, err := netip.ParseAddr(host); err == nil && addr.IsValid() {
		addr = addr.Unmap()
		if addr.Is4() {
			c.BindAddress4 = addr
		} else {
			c.BindAddress6 = addr
		}
	} else if addrs.IsDomainName(host) {
		c.BindAddress = host
	} else {
		return ex.New("bad host: ", host)
	}
	return nil
}

func (c *Conf) applyPort(port string) error {
	if numPort, err := strconv.ParseUint(port, 10, 16); err == nil {
		c.Port = uint16(numPort)
	} else if len(c.Network) == 1 {
		// when only one network(protocol) exist , lookup port from /etc/services(platform independent)
		lookupPort, err := net.LookupPort(c.Network[0].String(), port)
		if err != nil {
			return ex.Cause(err, "lookup port")
		}
		c.Port = uint16(lookupPort)
	} else {
		return ex.New("bad port: ", port)
	}
	return nil
}

func (c *Conf) Type() inbound.Type {
	return inbound.TypeDirect
}

func (c *Conf) String() string {
	return fmt.Sprint(*c)
}
