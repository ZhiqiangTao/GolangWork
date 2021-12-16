package stickpackage

import (
	"bytes"
	"encoding/binary"
	"net"
)

type LengthFieldBasedFrameCoder struct {
	//代表数据头，比如4，代表当前数据包的头4个字节代表的是数据段的总字节大小，读取的时候需要offset调
	header int32
	//其他字节，巴拉巴拉
}

func NewLengthFieldBasedFrameCoder() *LengthFieldBasedFrameCoder {
	return &LengthFieldBasedFrameCoder{header: 4}
}

// EncodeWrite client端用
func (f LengthFieldBasedFrameCoder) EncodeWrite(raw string, conn net.Conn) error {
	_raw := []byte(raw)
	h := make([]byte, 0, f.header)
	h = append(h, intToBytes(int32(len(_raw)))...) //长度数字，转成字节数组
	c := make([]byte, 0, len(_raw)+len(h))
	c = append(c, h...)
	c = append(c, _raw...)
	n, err := conn.Write(c) //把内容一口气发出去
	_ = n
	if err != nil {
		return err
	}
	return nil
}

// DecodeRead server端用
func (f LengthFieldBasedFrameCoder) DecodeRead(conn net.Conn) ([]byte, error) {
	h := make([]byte, f.header)
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
