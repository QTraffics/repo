package options

import (
	"github.com/qtraffics/qtfra/enhancements/slicelib"
	"github.com/qtraffics/qtfra/values"
)

type Option struct {
	DefaultResolver      string           `yaml:"defaultResolver"`
	DefaultDialInterface string           `yaml:"defaultDialInterface"`
	Log                  OptionLog        `yaml:"log"`
	Rules                []string         `yaml:"rules"`
	Resolvers            []OptionResolver `yaml:"resolvers"`
}

func Merge(oo ...Option) Option {
	if len(oo) == 0 {
		return Option{}
	}
	root := oo[0]
	for i := 1; i < len(oo); i++ {
		root.DefaultResolver = values.UseDefault(oo[i].DefaultResolver, root.DefaultResolver)
		root.DefaultDialInterface = values.UseDefault(oo[i].DefaultDialInterface, root.DefaultDialInterface)

		root.Log.Merge(oo[i].Log)
		root.Rules = append(root.Rules, oo[i].Rules...)
		root.Resolvers = append(root.Resolvers, oo[i].Resolvers...)
	}
	root.Rules = slicelib.Uniq(root.Rules)
	root.Resolvers = slicelib.UniqByLast(root.Resolvers, func(it OptionResolver) string {
		return it.Name
	})

	return root
}
