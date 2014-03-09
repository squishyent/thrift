package thrift

type TMultiplexedProtocolFactory struct {
	pf          TProtocolFactory
	serviceName string
}

func NewMultiplexedProtocolFactory(p TProtocolFactory, serviceName string) TMultiplexedProtocolFactory {
	return TMultiplexedProtocolFactory{p, serviceName}
}

func (f TMultiplexedProtocolFactory) GetProtocol(trans TTransport) TProtocol {
	p := NewMultiplexedProtocol(
		NewTProtocolDecorator(f.pf.GetProtocol(trans)),
		f.serviceName,
	)
	return &p
}
