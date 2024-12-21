package main

import "fmt"

func main() {
	edges := [][]int{{1, 2}, {1, 2}, {2, 3}, {4, 5}}

	for index, value := range edges {
		fmt.Println(index, value)
	}
}
