package stickpackage

import "net"

// ICoder 编码、解码器接口抽象
type ICoder interface {
	// EncodeWrite 客户端编码写方法
	EncodeWrite(raw string, conn net.Conn) error
	// DecodeRead 服务端解码读方法
	DecodeRead(conn net.Conn) ([]byte, error)
}
