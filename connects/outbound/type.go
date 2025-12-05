package outbound

import "strings"

type Type string

const (
	TypeDirect       Type = "direct"
	TypeLoadBalance  Type = "loadbalance"
	TypeFallback     Type = "fallback"
	TypeLatencyFirst Type = "latency-first"
)

func ParseTypeString(s string) Type {
	x := Type(strings.ToLower(s))
	switch x {
	case TypeDirect, TypeLoadBalance, TypeFallback, TypeLatencyFirst:
		return x
	default:
		return Type(s)
	}
}
