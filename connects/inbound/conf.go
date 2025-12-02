package inbound

type Conf interface {
	InBoundConf()

	String() string
}
