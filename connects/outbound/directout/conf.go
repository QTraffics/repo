package directout

func NewConfString(confString string) (Conf, error) {
	return Conf{}, nil
}

type Conf struct {
	TFO bool
}

func (c *Conf) OutboundConf() {}

func (c *Conf) String() string {
	return ""
}
