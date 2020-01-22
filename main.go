package main

import (
	"fmt"

	"github.com/niclic/golang-book/pkg/chapter1"
	"github.com/niclic/golang-book/pkg/chapter3"
)

func main() {
	fmt.Println("This is the main golang-book command.")

	chapter1.Hello()

	fmt.Println(chapter3.Add(2, 3))
	fmt.Println(chapter3.Multiply(32132, 42452))
}
