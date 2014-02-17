package thrift

/*
 * Adapted from Java: org.apache.thrift.protocol.TProtocolDecorator
 */

/**
 * <code>TProtocolDecorator</code> forwards all requests to an enclosed
 * <code>TProtocol</code> instance, providing a way to author concise
 * concrete decorator subclasses.  While it has no abstract methods, it
 * is marked abstract as a reminder that by itself, it does not modify
 * the behaviour of the enclosed <code>TProtocol</code>.
 *
 * <p>See p.175 of Design Patterns (by Gamma et al.)</p>
 *
 * @see org.apache.thrift.protocol.TMultiplexedProtocol
 */
type TProtocolDecorator struct {
	concreteProtocol TProtocol
}

func NewTProtocolDecorator(concreteProtocol TProtocol) TProtocolDecorator {
	return TProtocolDecorator{concreteProtocol}
}

func (p *TProtocolDecorator) WriteMessageBegin(name string, typeId TMessageType, seqid int32) error {
	return p.concreteProtocol.WriteMessageBegin(name, typeId, seqid)
}

func (p *TProtocolDecorator) WriteMessageEnd() error {
	return p.concreteProtocol.WriteMessageEnd()
}

func (p *TProtocolDecorator) WriteStructBegin(name string) error {
	return p.concreteProtocol.WriteStructBegin(name)
}

func (p *TProtocolDecorator) WriteStructEnd() error {
	return p.concreteProtocol.WriteStructEnd()
}

func (p *TProtocolDecorator) WriteFieldBegin(name string, typeId TType, id int16) error {
	return p.concreteProtocol.WriteFieldBegin(name, typeId, id)
}

func (p *TProtocolDecorator) WriteFieldEnd() error {
	return p.concreteProtocol.WriteFieldEnd()
}

func (p *TProtocolDecorator) WriteFieldStop() error {
	return p.concreteProtocol.WriteFieldStop()
}

func (p *TProtocolDecorator) WriteMapBegin(keyType TType, valueType TType, size int) error {
	return p.concreteProtocol.WriteMapBegin(keyType, valueType, size)
}

func (p *TProtocolDecorator) WriteMapEnd() error {
	return p.concreteProtocol.WriteMapEnd()
}

func (p *TProtocolDecorator) WriteListBegin(elemType TType, size int) error {
	return p.concreteProtocol.WriteListBegin(elemType, size)
}

func (p *TProtocolDecorator) WriteListEnd() error {
	return p.concreteProtocol.WriteListEnd()
}

func (p *TProtocolDecorator) WriteSetBegin(elemType TType, size int) error {
	return p.concreteProtocol.WriteSetBegin(elemType, size)
}

func (p *TProtocolDecorator) WriteSetEnd() error {
	return p.concreteProtocol.WriteSetEnd()
}

func (p *TProtocolDecorator) WriteBool(value bool) error {
	return p.concreteProtocol.WriteBool(value)
}

func (p *TProtocolDecorator) WriteByte(value byte) error {
	return p.concreteProtocol.WriteByte(value)
}

func (p *TProtocolDecorator) WriteI16(value int16) error {
	return p.concreteProtocol.WriteI16(value)
}

func (p *TProtocolDecorator) WriteI32(value int32) error {
	return p.concreteProtocol.WriteI32(value)
}

func (p *TProtocolDecorator) WriteI64(value int64) error {
	return p.concreteProtocol.WriteI64(value)
}

func (p *TProtocolDecorator) WriteDouble(value float64) error {
	return p.concreteProtocol.WriteDouble(value)
}

func (p *TProtocolDecorator) WriteString(value string) error {
	return p.concreteProtocol.WriteString(value)
}

func (p *TProtocolDecorator) WriteBinary(value []byte) error {
	return p.concreteProtocol.WriteBinary(value)
}

func (p *TProtocolDecorator) ReadMessageBegin() (name string, typeId TMessageType, seqid int32, err error) {
	return p.concreteProtocol.ReadMessageBegin()
}

func (p *TProtocolDecorator) ReadMessageEnd() error {
	return p.concreteProtocol.ReadMessageEnd()
}

func (p *TProtocolDecorator) ReadStructBegin() (name string, err error) {
	return p.concreteProtocol.ReadStructBegin()
}

func (p *TProtocolDecorator) ReadStructEnd() error {
	return p.concreteProtocol.ReadStructEnd()
}

func (p *TProtocolDecorator) ReadFieldBegin() (name string, typeId TType, id int16, err error) {
	return p.concreteProtocol.ReadFieldBegin()
}

func (p *TProtocolDecorator) ReadFieldEnd() error {
	return p.concreteProtocol.ReadFieldEnd()
}

func (p *TProtocolDecorator) ReadMapBegin() (keyType TType, valueType TType, size int, err error) {
	return p.concreteProtocol.ReadMapBegin()
}

func (p *TProtocolDecorator) ReadMapEnd() error {
	return p.concreteProtocol.ReadMapEnd()
}

func (p *TProtocolDecorator) ReadListBegin() (elemType TType, size int, err error) {
	return p.concreteProtocol.ReadListBegin()
}

func (p *TProtocolDecorator) ReadListEnd() error {
	return p.concreteProtocol.ReadListEnd()
}

func (p *TProtocolDecorator) ReadSetBegin() (elemType TType, size int, err error) {
	return p.concreteProtocol.ReadSetBegin()
}

func (p *TProtocolDecorator) ReadSetEnd() error {
	return p.concreteProtocol.ReadSetEnd()
}

func (p *TProtocolDecorator) ReadBool() (value bool, err error) {
	return p.concreteProtocol.ReadBool()
}

func (p *TProtocolDecorator) ReadByte() (value byte, err error) {
	return p.concreteProtocol.ReadByte()
}

func (p *TProtocolDecorator) ReadI16() (value int16, err error) {
	return p.concreteProtocol.ReadI16()
}

func (p *TProtocolDecorator) ReadI32() (value int32, err error) {
	return p.concreteProtocol.ReadI32()
}

func (p *TProtocolDecorator) ReadI64() (value int64, err error) {
	return p.concreteProtocol.ReadI64()
}

func (p *TProtocolDecorator) ReadDouble() (value float64, err error) {
	return p.concreteProtocol.ReadDouble()
}

func (p *TProtocolDecorator) ReadString() (value string, err error) {
	return p.concreteProtocol.ReadString()
}

func (p *TProtocolDecorator) ReadBinary() (value []byte, err error) {
	return p.concreteProtocol.ReadBinary()
}

func (p *TProtocolDecorator) Skip(fieldType TType) (err error) {
	return p.concreteProtocol.Skip(fieldType)
}

func (p *TProtocolDecorator) Flush() (err error) {
	return p.concreteProtocol.Flush()
}

func (p *TProtocolDecorator) Transport() TTransport {
	return p.concreteProtocol.Transport()
}
