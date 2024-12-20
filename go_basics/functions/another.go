package main

import "fmt"

func valueReturn(x, y, z int) (a, b, c int) {
	a = x + 10
	b = y + 10
	c = z + 10
	return
}

func main() {
	a, b := 10, 20
	fmt.Println(a, b)

	fmt.Println(valueReturn(10, 20, 39))
}
