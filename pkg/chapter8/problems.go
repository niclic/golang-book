package chapter8

// Zero sets the value of an integer to 0.
func Zero(iPtr *int) {
	*iPtr = 0 // dereference the pointer variable to change the underlying value
}

// One sets the value of an integer to 1.
func One(iPtr *int) {
	*iPtr = 1
}

// Square returns the square of two floating point numbers.
func Square(x *float64) {
	*x = *x * *x
}

// Swap swaps the values of two integers.
func Swap(x, y *int) {
	xPtr := *x // 1
	yPtr := *y // 2

	*x = yPtr
	*y = xPtr
	return
}
