package chapter4

import "fmt"

func FeetToMeters() {
	fmt.Println("Enter distance in ft: ")
	var ft float32
	fmt.Scanf("%f", &ft)

	m := (ft * 0.3048)

	fmt.Println("Distance in meters: ", m)
}
