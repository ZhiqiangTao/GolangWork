package main

import (
	"fmt"
	stickpackage "golangwork/week09/stickpackage/2-delimiter_based"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:7711")
	if err != nil {
		log.Fatalf("listen error:%v\n", err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error:%v\n", err)
			continue
		}
		go handleConn(conn)
	}
	fmt.Println("server closed!")
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	coder := stickpackage.NewDelimiterBasedCoder()
	for {
		read, err := coder.DecodeRead(conn)
		if err != nil {
			log.Printf("server read error\n")
			return
		}
		fmt.Printf("接收到的字节长度为：%v,\n数据为：%v \n\n", len(read), string(read))
	}
}
