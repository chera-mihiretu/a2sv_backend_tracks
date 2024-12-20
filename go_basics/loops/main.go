// go setup

package main

import "fmt"

func main() {

	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(i, sum)
	}

	// while loop

	for sum < 100 {
		sum += 10
		fmt.Println(sum)
	}
	count := 0
	for {
		count += 1
		// The if statement
		if count > 10 {
			break
		}
		fmt.Println(count)
	}
}
