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

func TestSetNoughtEmptyCell(t *testing.T) {
	b := NewBoard()
	err := b.Set(0, 0, "O")
	if err != nil {
		t.Fail()
	}
	v, err := b.Get(0, 0)
	if err != nil {
		t.Fail()
	}
	if v != "O" {
		t.Errorf("Expected O got %s", v)
	}
}

func TestSetOutOfBoundsError(t *testing.T) {
	b := NewBoard()
	var err error
	invalidCellTests := []struct {
		x int
		y int
		v string
	}{
		{42, 21, "X"},
		{1, -1, "O"},
	}
	for _, test := range invalidCellTests {
		err = b.Set(test.x, test.y, test.v)
		if err == nil {
			t.Errorf("No error found for cell (%d, %d)", test.x, test.y)
		}
	}
}

// TODO
// TestOverwriteCellFails
