package chapter7

// Average returns the average of a slice of integers.
func Average(ints ...int) int {
	var t int
	for _, i := range ints {
		t += i
	}
	return t / len(ints)
}

// Sum returns the sum of a slice of integers.
func Sum(ints ...int) int {
	var t int
	for _, i := range ints {
		t += i
	}
	return t
}
