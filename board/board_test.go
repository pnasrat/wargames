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
	var err error
	invalidCellTests := []struct {
		x int
		y int
	}{
		{42, 21},
		{1, -1},
	}
	for _, test := range invalidCellTests {
		_, err = b.Get(test.x, test.y)
		if err == nil {
			t.Errorf("No error found for cell (%d, %d)", test.x, test.y)
		}
	}
}
