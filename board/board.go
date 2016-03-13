package board

import "errors"

type Board struct {
	size  int
	board []string
}

func (b Board) Get(x, y int) (string, error) {
	if x < 0 || y < 0 || x >= b.size || y >= b.size {
		return "", errors.New("Out of board bounds")
	}
	return b.board[x+3*y], nil
}

func (b Board) Set(x, y int, v string) error {
	b.board[x+3*y] = v
	return nil
}

func NewBoard() *Board {
	return &Board{size: 3, board: make([]string, 3*3)}
}
