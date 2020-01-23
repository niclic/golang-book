package main

import (
	"fmt"

	"github.com/niclic/golang-book/pkg/chapter1"
	"github.com/niclic/golang-book/pkg/chapter3"
	"github.com/niclic/golang-book/pkg/chapter4"
	"github.com/niclic/golang-book/pkg/chapter5"
	"github.com/niclic/golang-book/pkg/chapter6"
)

func main() {
	fmt.Println("This is the main golang-book command.")

	chapter1.Hello()

	fmt.Println(chapter3.Add(2, 3))
	fmt.Println(chapter3.Multiply(32132, 42452))

	chapter4.FahrenheitToCelsius()
	chapter4.FeetToMeters()

	chapter5.DivisibleByThree()
	chapter5.FizzBuzz()
	chapter5.FizzBuzzRedux()

	nums := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	smallest := chapter6.FindSmallest(nums)
	fmt.Printf("The smallest number in %v is %d.", nums, smallest)
}
