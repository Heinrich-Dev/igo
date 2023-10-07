package main

import (
	"fmt"
)

// sets the board up to be sent between players
func setupBoard(boardSize int) []byte {
	return nil
}

// function takes in the bytes sent over the network and converts
// them into two integers representing the column and the row
// the player placed their piece on the other end
func ByteToBoard() (int, int) {
	return 0, 0
}

func BoardtoByte() {

}

func checkMove(move []byte, boardSize int) bool {
	if move[0] < 1 || move[0] > byte(boardSize) {
		fmt.Println("Column not in range of board")
		return false
	}
	if move[1] < 1 || move[1] > byte(boardSize) {
		fmt.Println("Row not in range of board")
		return false
	}
	return true
}

func GetUserInput(move []byte, boardSize int) {
	var moved bool = false
	for !moved {
		fmt.Println("Enter the column and row, space separated, that you would like to place your piece. (col row)")
		move[0] = 0
		move[1] = 0
		fmt.Scanln(&move[0], &move[1])
		if checkMove(move, boardSize) {
			moved = true
		}
	}
}
