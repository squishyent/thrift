package thrift

import (
	"net/http"
)

/*
 * An HTTP server for Thrift requests.
 *
 * Requests are sent as the body of a POST request sent to a particular URL.
 * If you wish to use this as a stand-alone HTTP server, start it with a call
 * to Serve() as normal. However, you can also use it as a handler in another
 * http server by calling Handle() as needed. This lets you support other URLs
 * on the same server.
 */

type THttpServer struct {
	addr              string
	cors              bool
	certFile, keyFile string
	LastError         error

	processorFactory       TProcessorFactory
	serverTransport        TServerTransport
	inputTransportFactory  TTransportFactory
	outputTransportFactory TTransportFactory
	inputProtocolFactory   TProtocolFactory
	outputProtocolFactory  TProtocolFactory
}

// Prepares a HTTP server.
func NewHttpServer(addr string, processorFactory TProcessorFactory, inputProtocolFactory TProtocolFactory, outputProtocolFactory TProtocolFactory) *THttpServer {
	return &THttpServer{
		addr:                   addr,
		processorFactory:       processorFactory,
		inputTransportFactory:  NewTTransportFactory(),
		outputTransportFactory: NewTTransportFactory(),
		inputProtocolFactory:   inputProtocolFactory,
		outputProtocolFactory:  outputProtocolFactory,
	}
}

// Prepares a HTTP server using TLS.
func NewHttpsServer(addr, certFile, keyFile string, processorFactory TProcessorFactory, inputProtocolFactory TProtocolFactory, outputProtocolFactory TProtocolFactory) *THttpServer {
	return &THttpServer{
		addr:                   addr,
		certFile:               certFile,
		keyFile:                keyFile,
		processorFactory:       processorFactory,
		inputTransportFactory:  NewTTransportFactory(),
		outputTransportFactory: NewTTransportFactory(),
		inputProtocolFactory:   inputProtocolFactory,
		outputProtocolFactory:  outputProtocolFactory,
	}
}

// Enable or disable CORS support
func (srv *THttpServer) SetCorsEnabled(enabled bool) {
	srv.cors = enabled
}

// Starts listening to the address and processing requests
func (srv *THttpServer) Serve() error {
	if srv.certFile != "" {
		return http.ListenAndServeTLS(srv.addr, srv.certFile, srv.keyFile, http.HandlerFunc(srv.Handle))
	} else {
		return http.ListenAndServe(srv.addr, http.HandlerFunc(srv.Handle))
	}
}

// Meant to stop the server, but this is optional and not implemented here
func (srv *THttpServer) Stop() error {
	return nil
}

// Handles a single HTTP request
func (srv *THttpServer) Handle(w http.ResponseWriter, req *http.Request) {

	// Handle CORS requests
	if req.Method == "OPTIONS" {
		if !srv.cors {
			w.WriteHeader(http.StatusForbidden)
			return
		} else {
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Methods", "POST")
			if v := req.Header.Get("Access-Control-Request-Headers"); v != "" {
				w.Header().Add("Access-Control-Allow-Headers", v)
			}
			return
		}
	}

	// Prepare the protocol stack
	client := &StreamTransport{
		Reader: req.Body,
		Writer: w,
	}
	processor := srv.processorFactory.GetProcessor(client)
	inputTransport := srv.inputTransportFactory.GetTransport(client)
	outputTransport := srv.outputTransportFactory.GetTransport(client)
	inputProtocol := srv.inputProtocolFactory.GetProtocol(inputTransport)
	outputProtocol := srv.outputProtocolFactory.GetProtocol(outputTransport)
	if inputTransport != nil {
		defer inputTransport.Close()
	}
	if outputTransport != nil {
		defer outputTransport.Close()
	}

	// Process the request
	_, srv.LastError = processor.Process(inputProtocol, outputProtocol)
	if err, ok := srv.LastError.(TTransportException); ok && err.TypeId() == END_OF_FILE {
		srv.LastError = nil
	}
}

func (srv *THttpServer) ProcessorFactory() TProcessorFactory {
	return srv.processorFactory
}

func (srv *THttpServer) ServerTransport() TServerTransport {
	return srv.serverTransport
}

func (srv *THttpServer) InputTransportFactory() TTransportFactory {
	return srv.inputTransportFactory
}

func (srv *THttpServer) OutputTransportFactory() TTransportFactory {
	return srv.outputTransportFactory
}

func (srv *THttpServer) InputProtocolFactory() TProtocolFactory {
	return srv.inputProtocolFactory
}

func (srv *THttpServer) OutputProtocolFactory() TProtocolFactory {
	return srv.outputProtocolFactory
}
