package chapter5

import (
	"fmt"
	"strconv"
)

func DivisibleByThree() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func FizzBuzz() {
	for i := 1; i <= 15; i++ {

		var s string

		if i%3 == 0 {
			s += "Fizz"
		}

		if i%5 == 0 {
			s += "Buzz"
		}

		if len(s) == 0 {
			s = strconv.Itoa(i)
		}

		fmt.Println(s)
	}
}

func FizzBuzzRedux() {
	for i := 1; i <= 15; i++ {
		var s string
		if i%3 == 0 {
			s = "Fizz"
		}
		if i%5 == 0 {
			s += "Buzz"
		}

		if len(s) == 0 {
			fmt.Println(i)
		} else {
			fmt.Println(s)
		}
	}

}
