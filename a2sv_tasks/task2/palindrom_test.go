package main

import (
	"testing"
)

func TestPalindrome(t *testing.T) {
	tests := map[string]bool{
		"madam": true,
		"hello": false,
	}

	for test, expcted := range tests {
		result := PalindromeCheck(test)
		if result != expcted {
			t.Errorf("Expected %v but got %v", expcted, result)
		}
	}

}
