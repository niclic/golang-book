package chapter7

// Half returns half the provided integer and also 'true' if that integer is even, otherwise 'false'.
func Half(i int) (int, string) {
	p := "false"
	if i%2 == 0 {
		p = "true"
	}
	return int(i / 2), p
}

// Sum returns the sum of a slice of integers.
func Sum(ints ...int) int {
	var t int
	for _, i := range ints {
		t += i
	}
	return t
}

// Average returns the average of a slice of integers.
func Average(ints ...int) int {
	var t int
	for _, i := range ints {
		t += i
	}
	return t / len(ints)
}
