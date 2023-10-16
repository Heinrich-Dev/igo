package main

import (
	"fmt"
	"strings"
)

const (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
)

func PrintBoard(board [][]byte, boardSize int) {
	fmt.Print("   ")
	for i := 0; i < boardSize; i++ {
		if i+1 < 10 {
			fmt.Printf("%d  ", i+1)
		} else {
			fmt.Printf("%d ", i+1)
		}
	}
	fmt.Println()
	for i := 0; i < boardSize; i++ {
		if i+1 < 10 {
			fmt.Printf("%d  ", i+1)
		} else {
			fmt.Printf("%d ", i+1)
		}
		for j := 0; j < boardSize-1; j++ {
			if board[i][j] == EMPTY {
				fmt.Print("*--")
			} else if board[i][j] == WHITE {
				fmt.Print("0--")
			} else {
				fmt.Print(colorRed + "0" + colorReset)
				fmt.Print("--")
			}
		}
		fmt.Printf("* %d\n", i+1)
		if i != boardSize-1 {
			fmt.Print("   ")
			for k := 0; k < boardSize-1; k++ {
				fmt.Print("|  ")
			}
			fmt.Println("|")
		}
	}
	fmt.Print("   ")
	for i := 0; i < boardSize; i++ {
		if i+1 < 10 {
			fmt.Printf("%d  ", i+1)
		} else {
			fmt.Printf("%d ", i+1)
		}
	}
	fmt.Println()
}

func checkMove(move []byte, board [][]byte, boardSize int) bool {
	if move[0] < 1 || move[0] > byte(boardSize) {
		fmt.Println("Column not in range of board.")
		return false
	}
	if move[1] < 1 || move[1] > byte(boardSize) {
		fmt.Println("Row not in range of board.")
		return false
	}
	if board[move[0]-1][move[1]-1] != 0 {
		fmt.Println("There is already a piece in that space.")
		return false
	}
	return true
}

// gets user input if they would like to move or pass, then
// checks if the given move is valid in the following ways:
//		ensures the placement of the piece is on the board
//		ensures the piece being placed is not surrounded (if no captures occur)
//		ensures ko rule is followed (no repeated board positions)
func GetUserInput(move []byte, board [][]byte, boardSize int, color int) {
	var response1, response2 string
	for {
		fmt.Println("Enter the column and row, space separated, that you would like to place your piece. (col row)")
		fmt.Scanln(&response1, &response2)
		if response1 == "p" {
			move[0] = 0
			move[1] = 0
			break
		}
		response := response1 + " " + response2
		responseReader := strings.NewReader(response)
		fmt.Fscan(responseReader, &move[0], &move[1])
		if !checkMove(move, board, boardSize) { // check if move is in bounds
			break
		}
		translatedMove := []byte{move[0] - 1, move[1] - 1} //translate into 0-indexed coords
		Capture(translatedMove, board, color)
		break // check if move captured any pieces and if piece is surrounded
	}
}

func PlacePiece(move []byte, board [][]byte, color int) {
	board[move[0]-1][move[1]-1] = byte(color)
}
