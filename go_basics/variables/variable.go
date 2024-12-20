package main

import "fmt"

var one, two bool
var int_one, int_two int
var char_one, char_two string
var byte_one, byte_two byte

func main() {
	fmt.Println(one, two)
	fmt.Println(int_one, int_two)
	fmt.Println(char_one, char_two)
	fmt.Println(byte_one, byte_two)

	// these are the variable types
	// integers
	var one int = 10
	var two int32 = 20
	var four int64 = 300

	fmt.Println(one, two, four)

	// unsigned integers
	var one_u uint = 10
	var two_u uint32 = 20
	var four_u uint64 = 300

	fmt.Println(one_u, two_u, four_u)

	// runes == int32

	var one_runes rune = 1
	fmt.Println(one_runes)

	// float numbers
	var one_f float32 = 1.1 // 6 - 7 decimal points
	var two_f float64 = 2.2 // 15 decimal points

	fmt.Println(one_f, two_f)

}
