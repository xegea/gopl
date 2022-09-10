package main

import (
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	// Close the listener when the application closes
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine
		go handleRequest(conn)
	}
}

// Handles incomming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incomming data
	buf := make([]byte, 1024)
	// Read the incomming connection into a buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	}
	// Send a response back to person contacting us.
	conn.Write([]byte("Message Received."))
	// Close the connection when you're done
	conn.Close()
}
