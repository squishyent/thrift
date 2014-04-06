package thrift

import (
	"fmt"
	"testing"
)

const TestMultiplexedProcessor_FN = "function_name"
const TestMultiplexedProcessor_IN_TYPE = CALL
const TestMultiplexedProcessor_OUT_TYPE = REPLY
const TestMultiplexedProcessor_SEQ = 1
const TestMultiplexedProcessor_IN_VALUE = int32(27)
const TestMultiplexedProcessor_OUT_VALUE = int32(32)

func NewTMultiplexedProcessor() *TMultiplexedProcessor {
	return &TMultiplexedProcessor{}
}

type DummyProcessor struct {
	t *testing.T
}

func (p *DummyProcessor) Process(in, out TProtocol) (bool, TException) {
	if name, typeId, seqid, e := in.ReadMessageBegin(); e != nil {
		p.t.Fatalf("DummyProcessor.Process() ReadMessageBegin() failed: %s", e.Error())
	} else if name != TestMultiplexedProcessor_FN {
		p.t.Fatalf("DummyProcessor.Process() expected message %q but got %q", TestMultiplexedProcessor_FN, name)
	} else if typeId != TestMultiplexedProcessor_IN_TYPE {
		p.t.Fatalf("DummyProcessor.Process() expected message %q but got %q", TestMultiplexedProcessor_IN_TYPE, typeId)
	} else if seqid != TestMultiplexedProcessor_SEQ {
		p.t.Fatalf("DummyProcessor.Process() expected message %q but got %q", TestMultiplexedProcessor_SEQ, seqid)
	}

	if v, e := in.ReadI32(); e != nil {
		p.t.Fatalf("DummyProcessor.Process() ReadI32() failed: %s", e.Error())
	} else if v != TestMultiplexedProcessor_IN_VALUE {
		p.t.Fatalf("DummyProcessor.Process() ReadI32() expected %v but got %v", TestMultiplexedProcessor_IN_VALUE, v)
	}

	if e := in.ReadMessageEnd(); e != nil {
		p.t.Fatalf("DummyProcessor.Process() ReadMessageEnd() failed: %s", e.Error())
	}

	if e := out.WriteMessageBegin(TestMultiplexedProcessor_FN, TestMultiplexedProcessor_OUT_TYPE, TestMultiplexedProcessor_SEQ); e != nil {
		p.t.Fatalf("DummyProcessor.Process() WriteMessageBegin() failed: %s", e.Error())
	} else if e := out.WriteI32(TestMultiplexedProcessor_OUT_VALUE); e != nil {
		p.t.Fatalf("DummyProcessor.Process() WriteI32() failed: %s", e.Error())
	} else if e := out.WriteMessageEnd(); e != nil {
		p.t.Fatalf("DummyProcessor.Process() WriteMessageEnd() failed: %s", e.Error())
	} else if e := out.Flush(); e != nil {
		p.t.Fatalf("DummyProcessor.Process() Flush() failed: %s", e.Error())
	}

	return true, nil
}

func TestMultiplexedProcessorProcess(t *testing.T) {
	var message = fmt.Sprintf("[\"%s:%s\",%d,%d,%v]",
		MULTIPLEXED_SERVICE_NAME,
		TestMultiplexedProcessor_FN,
		TestMultiplexedProcessor_IN_TYPE,
		TestMultiplexedProcessor_SEQ,
		TestMultiplexedProcessor_IN_VALUE,
	)
	var response = fmt.Sprintf("[\"%s\",%d,%d,%v]",
		TestMultiplexedProcessor_FN,
		TestMultiplexedProcessor_OUT_TYPE,
		TestMultiplexedProcessor_SEQ,
		TestMultiplexedProcessor_OUT_VALUE,
	)

	tIn := NewTMemoryReader(message)
	pIn := NewTSimpleJSONProtocol(tIn)

	tOut := NewTMemoryBuffer()
	pOut := NewTSimpleJSONProtocol(tOut)

	p := NewTMultiplexedProcessor()
	processor := &DummyProcessor{t}
	p.RegisterProcessor(MULTIPLEXED_SERVICE_NAME, processor)

	if succ, e := p.Process(pIn, pOut); !succ || e != nil {
		t.Fatalf("Process failed with error: %s", e.Error())
	}

	s := tOut.String()
	if s != response {
		t.Fatalf("Expected response %q but got %q", response, s)
	}
}
