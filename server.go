package main

import (
	"log"
	"net"
	"net/rpc"
	"time"
)

type Factorial int

func (f Factorial) Calculate(factor int, result *int) error {
	startedAt := time.Now()

	*result = 1
	for n := factor; n > 1; n-- {
		*result = (*result) * n
	}

	log.Println("Elapsed time: ", time.Since(startedAt))
	return nil
}

func serveConnection(conn net.Conn) {
	log.Println(conn.LocalAddr())
	rpc.ServeConn(conn)
}

func main() {
	rpc.Register(new(Factorial))
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error on listener: ", err)
	}

	log.Println("Server started")

	for {
		conn, err := listener.Accept()
		if err == nil {
			go serveConnection(conn)
		}
	}
}
