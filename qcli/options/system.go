package options

import "github.com/qtraffics/qtfra/values"

type OptionSystem struct {
	DefaultDNS           string `yaml:"defaultDns"`
	DefaultDialInterface string `yaml:"defaultDialInterface"`
}

func (o OptionSystem) Merge(o1 OptionSystem) OptionSystem {
	o.DefaultDNS = values.UseDefault(o1.DefaultDNS, o.DefaultDNS)
	o.DefaultDialInterface = values.UseDefault(o1.DefaultDialInterface, o.DefaultDialInterface)

	return o
}
