package main

import "fmt"

func main() {
	var array [10]int

	for i := 0; i < 10; i++ {
		array[i] = i
	}

	new_arr := array[4:6]

	new_arr[0] = 100

	fmt.Println(array)

	fmt.Println(array[3:])

	fmt.Println(array[:4])

	fmt.Println(cap(array[:]))

	var ar []int

	fmt.Println(len(ar))
	var val bool = ar == nil
	fmt.Println(val)
}
