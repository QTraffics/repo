package options

import (
	"github.com/qtraffics/qtfra/values"
)

type OptionLog struct {
	Disabled bool   `yaml:"disabled"`
	Level    string `yaml:"level"`
	Output   string `yaml:"output"`

	// set `none` to disable time output
	// default is time.RFC3339
	TimeFormat string `yaml:"timeFormat"`
}

func (o OptionLog) Merge(o1 OptionLog) OptionLog {
	o.Disabled = values.UseDefault(o1.Disabled, o.Disabled)
	o.Level = values.UseDefault(o1.Level, o.Level)
	o.Output = values.UseDefault(o1.Output, o.Output)
	o.TimeFormat = values.UseDefault(o1.TimeFormat, o.TimeFormat)
	return o
}
