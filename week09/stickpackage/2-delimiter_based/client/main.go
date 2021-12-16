package main

import (
	"fmt"
	stickpackage "golangwork/week09/stickpackage/2-delimiter_based"
	"log"
	"net"
	"strconv"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:7711")
	if err != nil {
		log.Fatalf("remote ip resolve fault!")
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatalf("remote ip dial fault!")
	}

	defer conn.Close()

	go func(conn net.Conn) {
		i := 0
		coder := stickpackage.NewDelimiterBasedCoder()
		for {
			words := "{\"Id\":" + strconv.Itoa(i) + ",\"Name\":\"golang\",\"Message\":\"message\"}"
			coder.EncodeWrite(words, conn)
			i++
		}
	}(conn)

	select {}
	fmt.Println("client closed!")
}
