package chapter7

func Average(ints ...int) int {
	var t int
	for _, i := range ints {
		t += i
	}
	return t / len(ints)
}
