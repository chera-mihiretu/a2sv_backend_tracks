package main

import "testing"

func TestFrequecy(t *testing.T) {
	tests := map[string]map[rune]int{
		"hello": {
			'h': 1,
			'e': 1,
			'l': 2,
			'o': 1,
		},
		"world": {
			'w': 1,
			'o': 1,
			'r': 1,
			'l': 1,
			'd': 1,
		},
	}

	for test, expected := range tests {
		result := FreqCount(test)
		for val, freq := range expected {
			if result[val] != freq {
				t.Errorf("Expected %v but got %v", expected, result)
				break
			}
		}
	}
}
