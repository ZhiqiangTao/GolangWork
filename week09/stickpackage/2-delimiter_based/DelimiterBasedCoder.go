package stickpackage

import (
	"net"
)

type DelimiterBasedCoder struct {
	fixChar string
}

func NewDelimiterBasedCoder() *DelimiterBasedCoder {
	return &DelimiterBasedCoder{fixChar: "\n"}
}

func (f DelimiterBasedCoder) EncodeWrite(raw string, conn net.Conn) error {
	_raw := []byte(raw)
	c := make([]byte, 0, len(_raw)+1)
	c = append(c, _raw...)
	c = append(c, []byte(f.fixChar)...)
	_, err := conn.Write(c)
	if err != nil {
		return err
	}
	return nil
}

func (f DelimiterBasedCoder) DecodeRead(conn net.Conn) ([]byte, error) {
	buffer := make([]byte, 1)
	res := make([]byte, 1)
	for {
		n, err := conn.Read(buffer)
		if n == 0 {
			break
		} else {
			if string(buffer) == f.fixChar {
				//匹配到特殊字符，就返回
				break
			}
			res = append(res, buffer...)
		}

		if err != nil {
			return nil, err
		}
	}

	return res, nil
}
