package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
		go HandleTCPConnection(c)
	}
}

func Echo(c net.Conn, str string, delay time.Duration) {
	_, err := fmt.Fprintln(c, "\t", strings.ToUpper(str))
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", str)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(str))

}

func HandleTCPConnection(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go Echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}
