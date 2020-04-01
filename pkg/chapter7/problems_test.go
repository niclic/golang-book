package chapter7

import (
	"testing"
)

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

func TestSum(t *testing.T) {
	var ints = []int{1, 2, 3, 4, 5}
	total := Sum(ints...)
	if total != 15 {
		t.Errorf("Sum should be %d, but was %d", 15, total)
	}
}
