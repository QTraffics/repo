package rulefile

import (
	"fmt"
	"strings"

	"github.com/qtraffics/qnetwork/addrs"
	"github.com/qtraffics/qtfra/ex"
	"github.com/qtraffics/repo/connects/outbound"
	"github.com/qtraffics/repo/connects/outbound/directout"
)

const itemDelimiter = ","

var (
	ErrEmptyRule   = ex.New("empty rule")
	ErrRequireMore = ex.New("require")
	ErrSyntax      = ex.New("syntax error")
)

// Rule
// <Listen> <Remotes> <OutboundType> <OutboundConf>
type Rule struct {
	Listen  string
	Remotes []addrs.Socksaddr

	OutboundType outbound.DialType
	OutboundConf []outbound.Conf
}

func ParseRule(ruleString string) (Rule, error) {
	if len(ruleString) == 0 || strings.HasPrefix(ruleString, "#") {
		return Rule{}, ErrEmptyRule
	}

	fields := strings.Fields(ruleString)
	if len(fields) < 2 {
		return Rule{}, errRequire(2)
	}

	var rule Rule
	for idx, field := range fields {
		if strings.HasPrefix(field, "#") {
			break
		}
		switch idx {
		case 0:
			rule.Listen = field
		case 1:
			remotes, err := parseRemotes(field)
			if err != nil {
				return Rule{}, errSyntax(err.Error())
			}
			rule.Remotes = remotes
		case 2:
			dialType := outbound.ParseDialType(field)
			rule.OutboundType = dialType
		case 3:
			switch rule.OutboundType {
			case outbound.TypeDirect, outbound.TypeNone:
				conf, err := directout.NewConfString(field)
				if err != nil {
					return Rule{}, errSyntax(err.Error())
				}
				rule.OutboundConf = append(rule.OutboundConf, &conf)
			case outbound.TypeLoadBalance:
				panic("implement me")
			case outbound.TypeFallback:
				panic("implement me")
			case outbound.TypeLatencyFirst:
				panic("implement me")
			default:
				panic("unexcepted")
			}
		default:
			break
		}
	}

	return rule, nil
}

func parseRemotes(remoteString string) (remotes []addrs.Socksaddr, err error) {
	items := strings.Split(remoteString, itemDelimiter)
	for _, item := range items {
		socksAddr := addrs.FromParseSocksaddr(item)
		if !socksAddr.Addr.IsValid() && !addrs.IsDomainName(socksAddr.Fqdn) {
			err = ex.New("wrong remote address")
			return
		} else if socksAddr.Port == 0 {
			err = ex.New("missing port in address")
			return
		}
		remotes = append(remotes, socksAddr)
	}
	return remotes, nil
}

func errRequire(atLeast int) error {
	return fmt.Errorf("%w at least %d fields", ErrRequireMore, atLeast)
}

func errSyntax(err string) error {
	return fmt.Errorf("%w : %s", ErrSyntax, err)
}
