package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	address := "localhost:8000"
	fmt.Println("Starting to listen on: ", address)
	socket, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := socket.Accept()
		fmt.Println("accept a client")
		if err != nil {
			log.Print(err)
			continue
		}
		// 加了一个go关键字，就可以接收多个客户端请求，否则只能处理一个请求
		go HandleTCPConnection(c)
	}
}

func HandleTCPConnection(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("03:04:05\n"))
		if err != nil {
			fmt.Println(err) // write tcp 127.0.0.1:8000->127.0.0.1:52383: write: broken pipe
			return
		}
		time.Sleep(time.Second)
	}
}
