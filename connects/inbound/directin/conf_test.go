package directin

import (
	"net/netip"
	"testing"

	"github.com/qtraffics/qnetwork/meta"
	"github.com/qtraffics/qtfra/ex"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testConfParseCase struct {
	Except *Conf
	Input  string

	Actually *Conf
	Err      error
}

func TestConfParse(t *testing.T) {
	protocolAll := meta.ProtocolList{meta.ProtocolTCP, meta.ProtocolUDP}

	cases := []testConfParseCase{
		{
			Except: &Conf{
				Network: protocolAll,
				Port:    8444,
			},
			Input: "8444",
		},
		{
			Except: &Conf{
				Network:       meta.ProtocolList{meta.ProtocolTCP},
				Port:          8444,
				BindInterface: "lo",
				EnableTFO:     true,
				ReuseAddr:     true,
			},
			Input: "tcp://lo:8444,reuseaddr,tfo",
		},
		{
			Except: &Conf{
				Network:      protocolAll,
				Port:         8444,
				BindAddress4: ex.Must0(netip.ParseAddr("1.1.1.1")),
				EnableTFO:    true,
			},
			Input: "tcp+udp://1.1.1.1:8444,tfo=true",
		},
		{
			Except: &Conf{
				Network:      meta.ProtocolList{meta.ProtocolUDP},
				Port:         8444,
				BindAddress6: ex.Must0(netip.ParseAddr("fe80::27b2:7ea4:f38c:d05c")),
				EnableTFO:    true,
				ReuseAddr:    true,
			},
			Input: "udp://[fe80::27b2:7ea4:f38c:d05c]:8444,tfo,reuseaddr",
		},
		{
			Except: &Conf{
				Network:     protocolAll,
				Port:        8444,
				BindAddress: "localhost",
			},
			Input: "localhost:8444",
		},
	}

	for _, v := range cases {
		v.Actually, v.Err = NewConfFromString(v.Input)
		require.Nil(t, v.Err)
		assert.Equal(t, *v.Except, *v.Actually)
	}
}
