package main

/*
Henry Boekhoff
10/1/23
igo.go - Creating the Go board game in Go. Two clients ssh'd on the same
machine can play with one another.
REFERENCES:
Inspired by Lachlan Imel's termChess: https://github.com/brochacho01/termChess
Creating and accepting TCP connections: https://pkg.go.dev/net
Colored terminal output: https://twin.sh/articles/35/how-to-add-colors-to-your-console-terminal-output-in-go
*/

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
	//"flag"
)

const (
	HOST  = "localhost"
	PORT  = "8080"
	EMPTY = 0
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

	//useFile := flag.Bool("f", false, "Using a file to setup board")
	//flag.Parse()

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
	var opponentColor int
	responded = false
	// Establishing board size
	initialMessage := []byte{0}

	if connectionType == 0 {
		fmt.Println("Client choosing board size...")
		connection.Read(initialMessage)
		boardSize = int(initialMessage[0])
		// Get random color and send to client
		rand.Seed(time.Now().UnixNano())
		color = rand.Intn(2) + 1
		initialMessage[0] = byte(color)
		connection.Write(initialMessage)
		if color == WHITE {
			opponentColor = RED
		} else {
			opponentColor = WHITE
		}
	} else {
		for {
			fmt.Println("Choose size of the board. (9, 19)")
			fmt.Println("9x9 board size recommended for beginnners. 19x19 board is standard.")
			fmt.Scanln(&boardSize)
			if boardSize == 9 || boardSize == 19 {
				break
			}
		}
		initialMessage[0] = byte(boardSize)
		connection.Write(initialMessage)
		connection.Read(initialMessage)
		opponentColor = int(initialMessage[0])
		if opponentColor == WHITE {
			color = RED
		} else {
			color = WHITE
		}
	}
	board := make([][]byte, boardSize, boardSize)
	for i := 0; i < boardSize; i++ {
		board[i] = make([]byte, boardSize)
	}
	PrintBoard(board, boardSize)

	move := make([]byte, 2)
	if color == RED {
		GetUserInput(move, board, boardSize, color)
		if move[0] != 255 {
			PlacePiece(move, board, color)
			fmt.Printf("Your move: %d %d\n", move[0]+1, move[1]+1)
		} else {
			fmt.Println("Skipped turn!")
		}
		connection.Write(move)
		PrintBoard(board, boardSize)
	}
	for {
		fmt.Println("Not your turn.")
		connection.Read(move)
		if move[0] != 255 {
			PlacePiece(move, board, opponentColor)
			fmt.Printf("Your opponent's move: %d %d\n", move[0]+1, move[1]+1)
		} else {
			fmt.Println("Opponent skipped turn!")
		}
		PrintBoard(board, boardSize)

		GetUserInput(move, board, boardSize, color)
		if move[0] != 255 {
			PlacePiece(move, board, color)
			fmt.Printf("Your move: %d %d\n", move[0]+1, move[1]+1)
		} else {
			fmt.Println("Skipped turn!")
		}
		connection.Write(move)
		PrintBoard(board, boardSize)
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
