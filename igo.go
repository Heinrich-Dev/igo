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
	"math/rand"
	"net"
	"os"
	"time"
)

const (
	HOST  = "localhost"
	PORT  = "8080"
	WHITE = 1
	RED   = 2
)

/* Main serves to establish the connection between client and server and
run the main game loop */
func main() {
	var response string
	var connectionType int // 0 for host, 1 for client
	var responded bool = false
	var connection net.Conn // interface for client or server, returned by helper functions
	// Setting up connection
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

	var boardSize int
	var color int
	responded = false
	// Establishing board size
	initialMessage := []byte{0}

	if connectionType == 0 {
		fmt.Println("Client choosing board size...")
		// TODO: See if there's a way to replace this ugliness
		// Get board size from client
		connection.Read(initialMessage)
		boardSize = int(initialMessage[0])
		// Get random color and send to client
		rand.Seed(time.Now().UnixNano())
		color = rand.Intn(2) + 1
		initialMessage[0] = byte(color)
		connection.Write(initialMessage)

	} else {
		for {
			fmt.Println("Choose size of the board. (9, 19)")
			fmt.Println("9x9 board size recommended for beginnners. 19x19 board is standard.")
			fmt.Scanln(&boardSize)
			if boardSize == 9 || boardSize == 19 {
				break
			}
		}
		// TODO: Write forces byte[], see if there's a way to send server just an int without needless conversion
		initialMessage[0] = byte(boardSize)
		connection.Write(initialMessage)
		// Get color
		connection.Read(initialMessage)
		color = int(initialMessage[0])
		if color == RED {
			color = WHITE
		} else {
			color = RED
		}
	}
	fmt.Printf("%d\n", color)
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
