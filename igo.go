package main

/*
Henry Boekhoff
10/1/23
igo.go - Creating the Go board game in Go. Two clients ssh'd on the same
machine can play with one another.
REFERENCES:
Inspired by Lachlan Imel's termChess: https://github.com/brochacho01/termChess
Creating and accepting TCP connections: https://pkg.go.dev/net
*/

import (
	"fmt"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "8080"
)

func main() {
	var response string
	var connectionType int // 0 for host, 1 for client
	var responded bool = false
	var connection net.Conn // interface for client or server, returned by helper functions
	for !responded {
		fmt.Println("Would you like to host or connect? (h/c)")
		fmt.Scanln(&response)
		if response == "h" {
			responded = true
			connectionType, connection = createServer()
		} else if response == "c" {
			responded = true
			connectionType, connection = join()
		} else {
			fmt.Println("Did not specify connection type!")
		}
	}
	if connectionType == -1 {
		fmt.Fprintf(os.Stderr, "Error while trying to establish connection between client and server\n")
	}
	if connectionType == 0 {
		connection.Write([]byte("TESTING"))
	} else {
		message := make([]byte, 1024)
		connection.Read(message)
		fmt.Printf("%s\n", message)
	}
}

func createServer() (int, net.Conn) {
	fmt.Println("Starting server on port 8080...")
	listener, err := net.Listen("tcp", HOST+":"+PORT)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Listen error\n")
		return -1, nil
	}
	conn, err := listener.Accept()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Accept error\n")
		return -1, nil
	}

	return 0, conn
}

func join() (int, net.Conn) {
	fmt.Println("Joining local server...")
	conn, err := net.Dial("tcp", HOST+":"+PORT)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Dial error\n")
		return -1, nil
	}
	return 1, conn
}
