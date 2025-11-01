package main

import (
	"fmt"
	"net"
	"os"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	// Close the listener when main exits
	defer l.Close()

	// Loop to continuously accept connections
	for {
		// Accept a new connection
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue // Continue to the next iteration of the loop
		}

		// Handle the connection in a new goroutine
		// This allows the server to accept multiple connections concurrently
		go handleConnection(conn)
	}
}

// handleConnection handles a single client connection
func handleConnection(conn net.Conn) {
	// Close the connection when this function returns
	defer conn.Close()

	// Create a buffer to read the request
	// 1024 bytes should be enough for the simple request in this stage
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from connection: ", err.Error())
		return
	}

	// Define the simple HTTP 200 OK response
	// \r\n is a Carriage Return + Line Feed (CRLF), which is the standard
	// line ending for HTTP protocols.
	// The final \r\n signifies the end of the headers section (which is empty).
	response := "HTTP/1.1 200 OK\r\n\r\n"

	// Write the response back to the client
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing response: ", err.Error())
	}
}
