package main

import "fmt"

func main() {
	var arr [][]int

	dx, dy := 8, 9

	for i := 0; i < dx; i++ {
		var slice []int
		arr = append(arr, slice)
		for j := 0; j < dy; j++ {
			arr[i] = append(arr[i], i)
		}
	}

	fmt.Println(arr)

}
