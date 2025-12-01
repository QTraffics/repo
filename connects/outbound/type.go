package outbound

import "strings"

type DialType string

const (
	TypeNone DialType = ""

	TypeDirect       DialType = "direct"
	TypeLoadBalance  DialType = "loadbalance"
	TypeFallback     DialType = "fallback"
	TypeLatencyFirst DialType = "latency-first"
)

func ParseDialType(s string) DialType {
	x := DialType(strings.ToLower(s))
	switch x {
	case TypeDirect, TypeLoadBalance, TypeFallback, TypeLatencyFirst, TypeNone:
		return x
	default:
		return TypeNone
	}
}
