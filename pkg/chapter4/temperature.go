package chapter4

import "fmt"

func FahrenheitToCelsius() {
	fmt.Println("Enter a temp in F: ")
	var f float64
	fmt.Scanf("%f", &f)

	c := (f - 32*5/9)
	fmt.Println("Temp in C is: ", c)
}
