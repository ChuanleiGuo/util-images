package main

import (
	"bufio"
	"log"
	"net"
	"strconv"
)

const port = 80

func main() {
	server, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if server == nil {
		log.Fatalf("start server err: %v\n", err)
	}
	conns := clientConns(server)
	for {
		go handleConn(<-conns)
	}
}

func clientConns(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	go func() {
		for {
			client, err := listener.Accept()
			if err != nil {
				log.Printf("accept err: %v\n", err)
				continue
			}
			log.Printf("%v <-> %v\n", client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()
	return ch
}

func handleConn(client net.Conn) {
	b := bufio.NewReader(client)
	for {
		line, err := b.ReadBytes('\n')
		if err != nil {
			break
		}
		client.Write(line)
	}
}
