package main

import "fmt"

// I pray that all of this typecasting does not cause any issues because
// I am only working with positive numbers. - Past Henry

func kill(row byte, col byte, board [][]byte, checked [][]byte, color int) {
	var oppositeColor int
	if color == WHITE {
		oppositeColor = RED
	} else {
		oppositeColor = WHITE
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if checked[i][j] == 1 && board[i][j] == byte(oppositeColor) {
				board[i][j] = EMPTY
			}
		}
	}
}

// returns true if given move does not result in the loss of the placed piece
func Capture(move []byte, board [][]byte, color int) bool {
	row := move[0]
	col := move[1]
	boardSize := len(board)
	checked := make([][]byte, boardSize, boardSize)
	for i := 0; i < boardSize; i++ {
		checked[i] = make([]byte, boardSize)
	}
	if row != 0 {
		if Surrounded(row-1, col, board, checked, color) {
			fmt.Println("Piece captured!")
			kill(row-1, col, board, checked, color)
		}
	}
	if row != byte(len(board))-1 {
		if Surrounded(row+1, col, board, checked, color) {
			fmt.Println("Piece captured!")
			kill(row+1, col, board, checked, color)
		}
	}
	if col != 0 {
		if Surrounded(row, col-1, board, checked, color) {
			fmt.Println("Piece captured!")
			kill(row, col-1, board, checked, color)
		}
	}
	if col != byte(len(board))-1 {
		if Surrounded(row, col+1, board, checked, color) {
			fmt.Println("Piece captured!")
			kill(row, col+1, board, checked, color)
		}
	}
	return !(Surrounded(row, col, board, checked, color))
}

// returns true if given piece is surrounded and should be killed
func Surrounded(row byte, col byte, board [][]byte, checked [][]byte, color int) bool {
	if surroundCheck(row, col, board, checked, color) == 0 {
		return true
	}
	return false
}

// returns an integer. If 0, that means the piece in question is completely surrounded.
// Any positive integer means the piece has at least one liberty free and does not need
// to be removed from the board.
// Total increases as more spaces are checked, a piece in the middle needs to be surrounded
// on all four liberties. A piece on the side needs to have 3 liberties surrounded and
// a piece in the corner needs only two.
func surroundCheck(row byte, col byte, board [][]byte, checked [][]byte, color int) int {
	var tally int
	total := 0
	if row != 0 {
		total++
		tally += surroundCheckHelper(row-1, col, board, checked, color)
	}
	if row != byte(len(board))-1 {
		total++
		tally += surroundCheckHelper(row+1, col, board, checked, color)
	}
	if col != 0 {
		total++
		tally += surroundCheckHelper(row, col-1, board, checked, color)
	}
	if col != byte(len(board))-1 {
		total++
		tally += surroundCheckHelper(row, col+1, board, checked, color)
	}
	return (total - tally)
}

// returns 1 if current piece is surrounded and 0 otherwise
func surroundCheckHelper(row byte, col byte, board [][]byte, checked [][]byte, color int) int {
	if checked[row][col] == 1 {
		return 1
	}
	checked[row][col] = 1
	curr := board[row][col]
	if curr == EMPTY {
		return 0
	}
	if color == WHITE && curr == RED {
		return 1
	}
	if color == WHITE && curr == WHITE {
		if surroundCheck(row, col, board, checked, color) == 4 {
			return 1
		}
	}
	if color == RED && curr == WHITE {
		return 1
	}
	if color == RED && curr == RED {
		if surroundCheck(row, col, board, checked, color) == 4 {
			return 1
		}
	}
	return 0
}
