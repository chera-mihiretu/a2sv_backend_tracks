package main

import "fmt"

func main() {
	var one int = 20
	var float_one float32 = float32(one)
	var unisigned uint16 = uint16(float_one)
	fmt.Println(float_one, unisigned)

	fmt.Printf("This float_one variable is %T \n", float_one)
}
