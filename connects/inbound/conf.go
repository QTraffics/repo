package inbound

type Conf interface {
	Type() Type

	String() string
}
