package main

import (
	"fmt"
)

const (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
)

func PrintBoard(board [][]byte, boardSize int) {
	fmt.Printf("  ")
	for i := 1; i <= boardSize; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	for i := 0; i < boardSize; i++ {
		fmt.Printf("%d ", i+1)
		for j := 0; j < boardSize-1; j++ {
			if board[i][j] == EMPTY {
				fmt.Printf(" -")
			} else if board[i][j] == WHITE {
				fmt.Printf("0-")
			} else {
				fmt.Printf(colorRed + "0" + colorReset)
				fmt.Printf("-")
			}
		}
		if board[i][boardSize-1] == EMPTY {
			fmt.Printf(" ")
		} else if board[i][boardSize-1] == WHITE {
			fmt.Printf("0")
		} else {
			fmt.Printf(colorRed + "0" + colorReset)
		}
		fmt.Println()
	}
	fmt.Printf("  ")
	for i := 1; i <= boardSize; i++ {
		fmt.Printf("%d ", i)
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

func GetUserInput(move []byte, board [][]byte, boardSize int) {
	var moved bool = false
	for !moved {
		fmt.Println("Enter the column and row, space separated, that you would like to place your piece. (col row)")
		move[0] = 0
		move[1] = 0
		fmt.Scanln(&move[0], &move[1])
		if checkMove(move, board, boardSize) {
			moved = true
		}
	}
}

func PlacePiece(move []byte, board [][]byte, color int) {
	board[move[0]-1][move[1]-1] = byte(color)
}
