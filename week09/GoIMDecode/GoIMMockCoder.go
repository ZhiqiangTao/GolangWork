package stickpackage

import (
	"bytes"
	"encoding/binary"
	"net"
)

type GoIMMockCoder struct {
	PackLen   int32 // package length， fix length
	HeaderLen int32 // header length
	Ver       int32 // protocol version
	Operation int32 // operation for request
	Seq       int32 // sequence number chosen by client
}

func NewGoIMMockCoder() *GoIMMockCoder {
	return &GoIMMockCoder{
		PackLen:   4,
		HeaderLen: 4,
		Ver:       4,
		Operation: 4,
		Seq:       4,
	}
}

func (coder GoIMMockCoder) CalcBodyLen(rawBodyLen int32) int32 {
	return rawBodyLen + coder.HeaderLen + coder.Ver + coder.Operation + coder.Seq
}

func (coder GoIMMockCoder) CalcPreLen() int32 {
	return coder.HeaderLen + coder.Ver + coder.Operation + coder.Seq
}

func build(fixLen, actualLen int32) []byte {
	res := make([]byte, 0, fixLen)
	b := intToBytes(actualLen)
	res = append(res, b...)
	return res
}

// EncodeWrite client端用 mock用！！！！
func (f GoIMMockCoder) EncodeWrite(header, ver, operation, seq, body string, conn net.Conn) error {
	_body := []byte(body)
	_total := f.CalcBodyLen(int32(len(_body))) //数据长度
	c := make([]byte, 0, _total+f.PackLen)     //整个容器

	//写PackLen
	c = append(c, build(f.PackLen, _total)...)
	//写HeaderLen
	c = append(c, build(f.HeaderLen, int32(len(header)))...)
	//写Ver
	c = append(c, build(f.Ver, int32(len(ver)))...)
	//写Operation
	c = append(c, build(f.Operation, int32(len(operation)))...)
	//写Seq
	c = append(c, build(f.Seq, int32(len(seq)))...)
	//写Body
	c = append(c, _body...)
	n, err := conn.Write(c)
	_ = n
	if err != nil {
		return err
	}
	return nil
}

// DecodeRead server端用
func (f GoIMMockCoder) DecodeRead(conn net.Conn) ([]byte, error) {
	h := make([]byte, f.PackLen)
	_, err := conn.Read(h)
	if err != nil {
		return nil, err
	}
	size := bytesToInt(h)
	res := make([]byte, 0, size)
	for size > 0 {
		c := make([]byte, size)
		n, err := conn.Read(c)
		if err != nil {
			return nil, err
		}

		res = append(res, c...)
		size = size - int32(n)
	}

	res = res[f.CalcPreLen():]
	return res, nil
}

func intToBytes(n int32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}

func bytesToInt(b []byte) int32 {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}
