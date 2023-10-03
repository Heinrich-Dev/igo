package main

/*
Henry Boekhoff
10/1/23
igo.go - Creating the Go board game in Go. Two clients ssh'd on the same
machine can play with one another.
REFERENCES:
Loosely Based on Lachlan Imel's termChess: https://github.com/brochacho01/termChess
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
	//var connectionType int // 0 for host, 1 for client
	var responded bool = false
	for !responded {
		fmt.Println("Would you like to host or connect? (h/c)")
		fmt.Scanf("%v", &response)
		if response[0] == 'h' {
			responded = true
			//connectionType = 0
			createServer()
		} else if response[0] == 'c' {
			responded = true
			//connectionType = 1
			join()
		} else {
			fmt.Println("Did not specify connection type!")
		}
	}

}

func createServer() (int, net.Conn) {
	fmt.Println("Starting server on port 8080...")
	listener, err := net.Listen("tcp", HOST+":"+PORT)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Listen error\n")
		return 1, nil
	}
	listener.Close()
	conn, err := listener.Accept()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Accept error\n")
		return 1, nil
	}
	numBytes, err := conn.Write([]byte("TESTING"))
	if numBytes == 0 {
		fmt.Fprintf(os.Stderr, "Read, somehow, read 0 bytes\n")
	}
	return 0, conn
}

func join() (int, net.Conn) {
	fmt.Println("Joining local server...")
	conn, err := net.Dial("tcp", HOST+":"+PORT)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Dial error\n")
		return 1, nil
	}

	message := make([]byte, 1024)
	numBytes, err := conn.Read(message)
	if err != nil || numBytes == 0 {
		fmt.Fprintf(os.Stderr, "Read failed\n")
	}

	return 0, conn
}
