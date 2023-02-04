package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	isPalindrome := isPalindrome(1221)

	fmt.Println(isPalindrome)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	original := x
	reversed := 0

	for x > 0 {
		reversed = (reversed * 10) + (x % 10)
		x /= 10
	}

	return original == reversed
}

func isPalindromeWithString(x int) bool {
	if x < 0 {
		return false
	}

	original := strconv.Itoa(x)
	reversed := ""

	originalChars := strings.Split(original, "")
	for i := len(originalChars) - 1; i >= 0; i-- {
		reversed += originalChars[i]
	}

	return original == reversed
}

func isPalindrome2(x int) bool {
	// Convert the integer to a string
	s := strconv.Itoa(x)

	// Compare the string with its reverse
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}