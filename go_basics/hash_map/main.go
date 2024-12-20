package main

import "fmt"

func main() {
	var mapping = make(map[int]int)

	for i := 0; i < 20; i++ {
		mapping[i] = 0
	}

	delete(mapping, 10)

	fmt.Println(mapping)
}
