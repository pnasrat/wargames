package board

import "errors"

type Board struct {
	size int
}

func (b Board) Get(x, y int) (string, error) {
	if x >= b.size || y >= b.size {
		return "", errors.New("Out of board bounds")
	}
	return "", nil
}

func NewBoard() *Board {
	return &Board{size: 3}
}
