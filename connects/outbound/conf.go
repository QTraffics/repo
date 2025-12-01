package outbound

type Conf interface {
	Apply(Outbound) error

	String() string
}
