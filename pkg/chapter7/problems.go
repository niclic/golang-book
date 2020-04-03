package chapter7

// Fib returns the Fibonnaci number for a given integer.
func Fib(i int) int {
	if i == 0 {
		return i
	}
	if i == 1 {
		return i
	}
	return Fib(i-1) + Fib(i-2)
}

// MakeOddGenerator is a generator that returns the next odd number each time it is called.
func MakeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i // this is an odd generator, so start at 1, not zero (even)
		i++     // increment by 1

		if i%2 == 0 {
			i++ // if i is even, increment by 1 again
		}
		return
	}
}

// FindGreatest finds the greatest integer in a list of integers.
func FindGreatest(nums ...int) int {
	g := 0
	for _, i := range nums {
		if i > g {
			g = i
		}
	}
	return g
}

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
