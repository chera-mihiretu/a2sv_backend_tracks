package main

import (
	"fmt"
	"math"
)

func powerLimit(x float64) bool {
	result := math.Pow(x, 4)
	if result < 100 {
		return true
	}
	return false

}
func main() {
	fmt.Println(powerLimit(2))
	fmt.Println(powerLimit(20))

	result := 10

	switch result {
	case 10:
		fmt.Println("The result is 10")
	case 20:
		fmt.Println("The result is 20")
	default:
		fmt.Println("The result is not 10 or 20")
	}
}
