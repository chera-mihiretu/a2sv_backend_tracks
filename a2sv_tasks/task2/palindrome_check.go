package main

func PalindromeCheck(word string) bool {
	palindrome := true
	for start, end := 0, len(word)-1; start < end; {
		if word[start] != word[end] {
			palindrome = false
			break
		}
		start++
		end--
	}
	return palindrome
}
