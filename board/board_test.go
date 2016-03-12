package board

import "testing"

func TestGetEmptyCell(t *testing.T) {
	b := NewBoard()
	v, err := b.Get(0, 0)
	if err != nil {
		t.Fail()
	}
	if v != "" {
		t.Fail()
	}
}

func TestGetOutOfBoundsError(t *testing.T) {
	b := NewBoard()
	_, err := b.Get(42, 21)
	if err == nil {
		t.Fail()
	}
}
