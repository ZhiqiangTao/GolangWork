package main

import (
	"fmt"
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
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024) //容器1024个字节
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("server read error\n")
			return
		}
		fmt.Printf("接收到的字节长度为：%v, 数据为：%v \n\n", n, string(buffer))
	}
}
