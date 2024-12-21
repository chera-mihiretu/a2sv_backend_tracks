package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	parts := strings.Split(s, " ")
	count := make(map[string]int)
	for _, value := range parts {
		_, k := count[value]
		if !k {
			count[value] = 0
		}
		count[value] += 1
	}
	return count
}

func main() {
	fmt.Println(WordCount("I am Chera Mihiretu Go!"))
}
