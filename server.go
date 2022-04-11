package main

import (
	"log"
	"net"
	"net/rpc"
	"time"
)

// Declares a type whose methods will be made available by rpc
type Factorial int

// Declares a method for Factorial that calculates the factorial of factor value and store it using result pointer
func (f Factorial) Calculate(factor int, result *int) error {
	startedAt := time.Now()

	*result = 1
	for n := factor; n > 1; n-- {
		*result = (*result) * n
	}

	log.Println("Elapsed time: ", time.Since(startedAt))
	return nil
}

// Declares a function that prints the Address of the client connection and serves the rpc response
func serveConnection(conn net.Conn) {
	log.Println(conn.LocalAddr())
	rpc.ServeConn(conn)
}

func main() {
	rpc.Register(new(Factorial))
	rpc.HandleHTTP()

	// Declares a listener which will receive the client connections at port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error on listener: ", err)
	}

	log.Println("Server started")
	for {
		// Get a single connection to process
		conn, err := listener.Accept()
		if err == nil {
			// Work on a connection asynchronously
			go serveConnection(conn)
		}
	}
}
