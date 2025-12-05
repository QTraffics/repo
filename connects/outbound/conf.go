package outbound

type Conf interface {
	Type() Type

	String() string
}
