package godataserializer

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type DataSerializer struct {
	Order binary.ByteOrder
	Buffer bytes.Buffer
}

func NewDataSerializer(order binary.ByteOrder, data ...[]byte) *DataSerializer {
	s := &DataSerializer{
		Order: order,
	}

	if len(data) > 0 && len(data[0]) > 0 {
		s.Buffer.Write(data[0])
	}

	return s
}

func (s *DataSerializer) Bytes() []byte {
	return s.Buffer.Bytes()
}

func (s *DataSerializer) Next(n int) []byte {
	return s.Buffer.Next(n)
}

func (s *DataSerializer) Reset() {
	s.Buffer.Reset()
}

func (s *DataSerializer) Read(v interface{}) {
	if err := binary.Read(&s.Buffer, s.Order, v); err != nil {
		fmt.Println("DataSerializer read error: " + err.Error())
	}
}

func (s *DataSerializer) Write(v interface{}) {
	if err := binary.Write(&s.Buffer, s.Order, v); err != nil {
		fmt.Println("DataSerializer write error: " + err.Error())
	}
}

func (s *DataSerializer) ReadByte() (v byte) {
	s.Read(&v)
	return
}

func (s *DataSerializer) WriteByte(v byte) {
	s.Write(v)
}

func (s *DataSerializer) ReadInt8() (v int8) {
	s.Read(&v)
	return
}

func (s *DataSerializer) WriteInt8(v int8) {
	s.Write(v)
}

func (s *DataSerializer) ReadUint8() (v uint8) {
	s.Read(&v)
	return
}

func (s *DataSerializer) WriteUint8(v uint8) {
	s.Write(v)
}

func (s *DataSerializer) ReadInt16() (v int16) {
	s.Read(&v)
	return
}

func (s *DataSerializer) WriteInt16(v int16) {
	s.Write(v)
}

func (s *DataSerializer) ReadUint16() (v uint16) {
s.Read(&v)
return
}

func (s *DataSerializer) WriteUint16(v uint16) {
	s.Write(v)
}

func (s *DataSerializer) ReadInt32() (v int32) {
	s.Read(&v)
	return
}

func (s *DataSerializer) WriteInt32(v int32) {
	s.Write(v)
}

func (s *DataSerializer) ReadUint32() (v uint32) {
	s.Read(&v)
	return
}

func (s *DataSerializer) WriteUint32(v uint32) {
	s.Write(v)
}

func (s *DataSerializer) ReadInt64() (v int64) {
	s.Read(&v)
	return
}

func (s *DataSerializer) WriteInt64(v int64) {
	s.Write(v)
}

func (s *DataSerializer) ReadUint64() (v uint64) {
	s.Read(&v)
	return
}

func (s *DataSerializer) WriteUint64(v uint64) {
	s.Write(v)
}

func (s *DataSerializer) ReadFloat32() (v float32) {
	s.Read(&v)
	return
}

func (s *DataSerializer) WriteFloat32(v float32) {
	s.Write(v)
}

func (s *DataSerializer) ReadFloat64() (v float64) {
	s.Read(&v)
	return
}

func (s *DataSerializer) WriteFloat64(v float64) {
	s.Write(v)
}

func (s *DataSerializer) ReadBool() (v bool) {
	s.Read(&v)
	return
}

func (s *DataSerializer) WriteBool(v bool) {
	s.Write(v)
}

func (s *DataSerializer) ReadString() string {
	return string(s.Next(int(s.ReadInt16())))
}

func (s *DataSerializer) WriteString(v string) {
	s.WriteInt16(int16(len(v)))
	s.Write([]byte(v))
}