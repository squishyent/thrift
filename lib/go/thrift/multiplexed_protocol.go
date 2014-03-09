package thrift

const (
	SEPARATOR = ":"
)

/**
 * <code>TMultiplexedProtocol</code> is a protocol-independent concrete decorator
 * that allows a Thrift client to communicate with a multiplexing Thrift server,
 * by prepending the service name to the function name during function calls.
 *
 * <p>NOTE: THIS IS NOT USED BY SERVERS.  On the server, use {@link org.apache.thrift.TMultiplexedProcessor TMultiplexedProcessor} to handle requests
 * from a multiplexing client.
 *
 * <p>This Java example uses a single socket transport to invoke two services:
 *
 * <blockquote><code>
 *     TSocket transport = new TSocket("localhost", 9090);<br/>
 *     transport.open();<br/>
 *<br/>
 *     TBinaryProtocol protocol = new TBinaryProtocol(transport);<br/>
 *<br/>
 *     TMultiplexedProtocol mp = new TMultiplexedProtocol(protocol, "Calculator");<br/>
 *     Calculator.Client service = new Calculator.Client(mp);<br/>
 *<br/>
 *     TMultiplexedProtocol mp2 = new TMultiplexedProtocol(protocol, "WeatherReport");<br/>
 *     WeatherReport.Client service2 = new WeatherReport.Client(mp2);<br/>
 *<br/>
 *     System.out.println(service.add(2,2));<br/>
 *     System.out.println(service2.getTemperature());<br/>
 * </code></blockquote>
 *
 * @see org.apache.thrift.protocol.TProtocolDecorator
 */
type TMultiplexedProtocol struct {
	TProtocolDecorator
	serviceName string
}

func NewMultiplexedProtocol(p TProtocolDecorator, serviceName string) TMultiplexedProtocol {
	return TMultiplexedProtocol{p, serviceName}
}

/**
 * Prepends the service name to the function name, separated by TMultiplexedProtocol.SEPARATOR.
 *
 * @param tMessage The original message.
 * @throws TException Passed through from wrapped <code>TProtocol</code> instance.
 */
func (p *TMultiplexedProtocol) WriteMessageBegin(name string, typeId TMessageType, seqid int32) error {
	if typeId == CALL || typeId == ONEWAY {
		return p.concreteProtocol.WriteMessageBegin(
			p.serviceName+SEPARATOR+name,
			typeId,
			seqid,
		)
	} else {
		return p.concreteProtocol.WriteMessageBegin(name, typeId, seqid)
	}
}
