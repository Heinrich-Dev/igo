package main

import (
	"fmt"
)

// sets the board up to be sent between players
func setupBoard(boardSize int) []byte {
	return nil
}

func PrintBoard(board [][]byte) {

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
