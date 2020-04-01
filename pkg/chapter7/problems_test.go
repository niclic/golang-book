package chapter7

import (
	"testing"
)

func TestHalf(t *testing.T) {
	cases := []struct {
		num    int
		half   int
		isEven string
	}{
		{0, 0, "true"},
		{1, 0, "false"},
		{2, 1, "true"},
		{3, 1, "false"},
		{4, 2, "true"},
		{5, 2, "false"},
		{10, 5, "true"},
		{15, 7, "false"},
	}

	for _, c := range cases {
		res, isEven := Half(c.num)

		if res != c.half {
			t.Error("Input/Output mismatch", c.num, c.half, res)
		}
		if isEven != c.isEven {
			t.Error("Parity is incorrect", c.num, c.isEven, isEven)
		}
	}
}

func TestSum(t *testing.T) {
	var ints = []int{1, 2, 3, 4, 5}
	total := Sum(ints...)
	if total != 15 {
		t.Errorf("Sum should be %d, but was %d", 15, total)
	}
}

func TestAverage(t *testing.T) {
	cases := []struct {
		nums []int
		avg  int
	}{
		{[]int{1, 2, 3}, 2},
		{[]int{10, 45, 3, 8, 5, 17}, 14},
		{[]int{5, 13, 22}, 13},
	}

	for _, c := range cases {
		avg := Average(c.nums...)

		if avg != c.avg {
			t.Error("The average is incorrect", c.nums, avg)
		}
	}
}
