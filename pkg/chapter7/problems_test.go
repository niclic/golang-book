package chapter7

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	cases := []struct {
		seq int
		num int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{33, 3524578},
	}

	for _, c := range cases {
		num := Fib(c.seq)
		if num != c.num {
			t.Errorf("Fibonacci number %d: expected %d, but got %d instead.", c.seq, c.num, num)
		}
	}
}

func TestMakeOddGenerator(t *testing.T) {
	makeOdd := MakeOddGenerator()
	length := 5
	var nums = make([]uint, length)
	for i := 0; i < length; i++ {
		nums = append(nums, makeOdd())
	}

	expected := []uint{1, 3, 5, 7, 9}
	if reflect.DeepEqual(nums, expected) {
		fmt.Printf("makeOddGenerator is not working. Expected %d, got %d", length, len(nums))
	}
}

func TestFindGreatest(t *testing.T) {
	nums := []int{1, 24, 12, 66, 4, 198, 44, 3}
	expected := 198
	actual := FindGreatest(nums...)
	if actual != expected {
		t.Errorf("Expected %d to be the greatest, but it was %d", expected, actual)
	}
}

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
