package options

import (
	"github.com/qtraffics/qtfra/enhancements/slicelib"
)

type Option struct {
	System    OptionSystem `yaml:"system"`
	Log       OptionLog    `yaml:"log"`
	Rules     []string     `yaml:"rules"`
	RuleFiles []string     `yaml:"ruleFiles"`
}

func Merge(oo ...Option) Option {
	if len(oo) == 0 {
		return Option{}
	}
	root := oo[0]
	for i := 1; i < len(oo); i++ {
		root.System.Merge(oo[i].System)
		root.Log.Merge(oo[i].Log)

		root.Rules = append(root.Rules, oo[i].Rules...)
		root.RuleFiles = append(root.RuleFiles, oo[i].RuleFiles...)
	}
	root.Rules = slicelib.Uniq(root.Rules)
	root.RuleFiles = slicelib.Uniq(root.RuleFiles)

	return root
}
