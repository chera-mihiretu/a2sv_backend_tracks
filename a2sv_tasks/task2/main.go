package main

import "fmt"

func main() {
	fmt.Println("Enter The Word you want to count the caracters :")
	var word string

	fmt.Scan(&word)

	var count map[rune]int = FreqCount(word)

	for val, fre := range count {
		fmt.Printf("%c : %d\n", val, fre)
	}
	palindrome := PalindromeCheck(word)
	if palindrome {
		fmt.Println("The word is a palindrome")
	} else {
		fmt.Println("The word is not a palindrome")
	}
}
