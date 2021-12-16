package stickpackage

import (
	"log"
	"net"
)

type FixedLengthCoder struct {
	//固定的字节长度
	fixLen int
}

func NewFixedLengthCoder() *FixedLengthCoder {
	l := 256
	return &FixedLengthCoder{fixLen: l}
}

func (f FixedLengthCoder) EncodeWrite(raw string, conn net.Conn) error {
	bs := []byte(raw)
	if len(bs) > f.fixLen {
		log.Fatalf("fixed length 模式下，上传数据不能超过固定长度，请适当调整固定长度")
	}
	offset := 0
	dataLeft := len(bs)
	for dataLeft > 0 {
		container := make([]byte, 0, f.fixLen) //固定长度
		container = append(container, bs...)
		dif := f.fixLen - len(bs)
		if dif > 0 {
			container = append(container, make([]byte, dif, dif)...) //不足补0
		}
		sent, err := conn.Write(container)
		if err != nil {
			return err
		}
		bs = bs[offset:min(f.fixLen, len(bs))]
		offset += sent
		dataLeft -= sent
	}
	return nil
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func (f FixedLengthCoder) DecodeRead(conn net.Conn) ([]byte, error) {
	buffer := make([]byte, f.fixLen)
	_, err := conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	//buffer = bytes.Trim(buffer, "\x00")
	return buffer, nil
}
