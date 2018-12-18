package main

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(1, 2)
	if sum != 3 {
		t.Error("Add 1 and 2 result is not 3!")
	}
}

func TestAdd2(t *testing.T) {
	sum := Add(0, 2)
	if sum != 2 {
		t.Error("Add 0 and 2 result is not 2!")
	}
}

func TestAddTable(t *testing.T) {
	var tests = []struct {
		arg1 int
		arg2 int
		sum  int
	}{
		{1, 1, 2},
		{-1, -1, -2},
		{-1, 1, 0},
	}

	for _, test := range tests {
		sum := Add(test.arg1, test.arg2)
		if sum != test.sum {
			t.Errorf("Add %v and %v result is not %v !", test.arg1, test.arg2, test.sum)
		}
	}
}
