package chapter8

import "testing"

func TestZero(t *testing.T) {
	i := 5
	Zero(&i)
	if i != 0 {
		t.Errorf("Zero is not working. Expected 0, got %d", i)
	}
}

func TestOne(t *testing.T) {
	iPtr := new(int)
	One(iPtr)
	if *iPtr != 1 {
		t.Errorf("Zero is not working with new. Expected 1, got %d", *iPtr)
	}
}

func TestSquare(t *testing.T) {
	h := 1.5
	expected := 2.25
	Square(&h)
	if h != expected {
		t.Errorf("Square is not working. Expected %f, got %f", expected, h)
	}
}

func TestSwap(t *testing.T) {
	x := 1
	y := 2
	Swap(&x, &y)
	if x != 2 && y != 1 {
		t.Errorf("Swap is not working. Expected x = 2, but is %d, and y = 1, but is %d.", x, y)
	}
}
