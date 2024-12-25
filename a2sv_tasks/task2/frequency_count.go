package main

func FreqCount(input string) map[rune]int {
	frequency := make(map[rune]int)
	for _, char := range input {
		if _, exist := frequency[char]; !exist {
			frequency[char] = 0
		}
		frequency[char]++
	}

	return frequency
}
