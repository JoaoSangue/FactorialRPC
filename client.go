//
// Use `go run client.go` to start application
//
package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	// Connect to server at "localhost:8080" using TCP
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}

	var value, result int
	fmt.Println("Enter a number: ")
	fmt.Scanf("%d", &value)

	// Ask server to execute "Factorial.Calculate(value, &result)"
	err = client.Call("Factorial.Calculate", value, &result)
	if err != nil {
		log.Fatal(fmt.Sprintf("Couldn't calculate factorial of %d. ", value), err)
	}

	fmt.Println("Result: ", result)
}
