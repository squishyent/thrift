package thrift

/*
 * Adapted from Java: org.apache.thrift.TMultiplexedProcessor
 */

import (
	"fmt"
	"strings"
)

/**
 *  Our goal was to work with any protocol.  In order to do that, we needed
 *  to allow them to call ReadMessageBegin() and get a TMessage in exactly
 *  the standard format, without the service name prepended to TMessage.name.
 */
type StoredMessageProtocol struct {
	TProtocolDecorator
	name   string
	typeId TMessageType
	seqid  int32
	err    error
}

func (p *StoredMessageProtocol) ReadMessageBegin() (name string, typeId TMessageType, seqid int32, err error) {
	return p.name, p.typeId, p.seqid, p.err
}

/**
 * <code>TMultiplexedProcessor</code> is a <code>TProcessor</code> allowing
 * a single <code>TServer</code> to provide multiple services.
 *
 * <p>To do so, you instantiate the processor and then register additional
 * processors with it, as shown in the following example:</p>
 *
 * <blockquote><code>
 *     TMultiplexedProcessor processor = new TMultiplexedProcessor();
 *
 *     processor.registerProcessor(
 *         "Calculator",
 *         new Calculator.Processor(new CalculatorHandler()));
 *
 *     processor.registerProcessor(
 *         "WeatherReport",
 *         new WeatherReport.Processor(new WeatherReportHandler()));
 *
 *     TServerTransport t = new TServerSocket(9090);
 *     TSimpleServer server = new TSimpleServer(processor, t);
 *
 *     server.serve();
 * </code></blockquote>
 */
type TMultiplexedProcessor struct {
	processors map[string]TProcessor
}

/**
 * 'Register' a service with this <code>TMultiplexedProcessor</code>.  This
 * allows us to broker requests to individual services by using the service
 * name to select them at request time.
 *
 * @param serviceName Name of a service, has to be identical to the name
 * declared in the Thrift IDL, e.g. "WeatherReport".
 * @param processor Implementation of a service, ususally referred to
 * as "handlers", e.g. WeatherReportHandler implementing WeatherReport.Iface.
 */
func (p *TMultiplexedProcessor) RegisterProcessor(serviceName string, processor TProcessor) {
	if p.processors == nil {
		p.processors = make(map[string]TProcessor)
	}
	p.processors[serviceName] = processor
}

/**
 * This implementation of <code>process</code> performs the following steps:
 *
 * <ol>
 *     <li>Read the beginning of the message.</li>
 *     <li>Extract the service name from the message.</li>
 *     <li>Using the service name to locate the appropriate processor.</li>
 *     <li>Dispatch to the processor, with a decorated instance of TProtocol
 *         that allows ReadMessageBegin() to return the original TMessage.</li>
 * </ol>
 *
 * @throws TException If the message type is not CALL or ONEWAY, if
 * the service name was not found in the message, or if the service
 * name was not found in the service map.  You called {@link #registerProcessor(String, TProcessor) registerProcessor}
 * during initialization, right? :)
 */
func (p *TMultiplexedProcessor) Process(in, out TProtocol) (bool, TException) {
	/*
	   Use the actual underlying protocol (e.g. TBinaryProtocol) to read the
	   message header.  This pulls the message "off the wire", which we'll
	   deal with at the end of this method.
	*/
	name, typeId, seqid, err := in.ReadMessageBegin()
	if err != nil {
		return false, err
	}

	if typeId != CALL && typeId != ONEWAY {
		// TODO Apache Guys - Can the server ever get an EXCEPTION or REPLY?
		// TODO Should we check for this here?
		return false, TException(fmt.Errorf("This should not have happened!?"))
	}

	// Extract the service name
	index := strings.Index(name, SEPARATOR)
	if index < 0 || index >= len(name)-1 {
		return false, TException(fmt.Errorf(
			"Service name not found in message name: " + name + ".  Did you " +
				"forget to use a TMultiplexProtocol in your client?"))
	}

	// Create a new TMessage, something that can be consumed by any TProtocol
	serviceName := name[:index]
	actualProcessor, ok := p.processors[serviceName]
	if !ok {
		return false, TException(fmt.Errorf(
			"Service name not found: " + serviceName + ".  Did you forget " +
				"to call registerProcessor()?"))
	}

	// Remove the service name from the message name
	standardName := name[index+1:]

	// Dispatch processing to the stored processor
	return actualProcessor.Process(
		&StoredMessageProtocol{
			NewTProtocolDecorator(in),
			standardName,
			typeId,
			seqid,
			err,
		}, out)
}
