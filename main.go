package main

import (
	"flag"
	"fmt"

	"github.com/niclic/golang-book/pkg/chapter1"
	"github.com/niclic/golang-book/pkg/chapter3"
	"github.com/niclic/golang-book/pkg/chapter4"
	"github.com/niclic/golang-book/pkg/chapter5"
	"github.com/niclic/golang-book/pkg/chapter6"
	"github.com/niclic/golang-book/pkg/chapter7"
)

func ch1() {
	fmt.Println("Chapter 1:")
	chapter1.Hello()
}

func ch2() {
}

func ch3() {
	fmt.Println("Chapter 3:")
	fmt.Println(chapter3.Add(2, 3))
	fmt.Println()
	fmt.Println(chapter3.Multiply(32132, 42452))
}

func ch4() {
	fmt.Println("Chapter 4:")
	chapter4.FahrenheitToCelsius()
	fmt.Println()
	chapter4.FeetToMeters()
}

func ch5() {
	fmt.Println("Chapter 5:")
	chapter5.DivisibleByThree()
	fmt.Println()
	chapter5.FizzBuzz()
	fmt.Println()
	chapter5.FizzBuzzRedux()
}

func ch6() {
	fmt.Println("Chapter 6:")

	sl := make([]int, 3, 9)
	fmt.Println(len(sl))

	fmt.Println()

	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x[2:5])
	fmt.Println()

	nums := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	smallest := chapter6.FindSmallest(nums)
	fmt.Printf("The smallest number in %v is %d.\n", nums, smallest)
}

func ch7() {
	fmt.Println("Chapter 7:")

	nums := []int{45, 11, 100, 5, 786, 33, 1, 4}
	average := chapter7.Average(nums...)
	fmt.Printf("The average of %v is %d\n", nums, average)

	sum := chapter7.Sum(nums...)
	fmt.Printf("The sum of %v is %d\n", nums, sum)
}

func main() {
	fmt.Println("This is the main golang-book command.")
	fmt.Println("Enter a chapter number to run those exercises: e.g. golang-book -ch=1")
	fmt.Println()

	chapter := flag.Int("ch", 1, "Chapter number")
	flag.Parse()

	switch *chapter {
	case 1:
		ch1()
	case 2:
		ch2()
	case 3:
		ch3()
	case 4:
		ch4()
	case 5:
		ch5()
	case 6:
		ch6()
	case 7:
		ch7()
	default:
		fmt.Println("Please enter a Chapter number.")
	}
}
