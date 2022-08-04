package p125

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	n := len(s)
	left, right := 0, n-1

	for left <= right {
		for left <= right && !access(s[left]) {
			left++
		}
		for left <= right && !access(s[right]) {
			right--
		}
		if left >= right {
			return true
		}
		if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func access(c uint8) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	if c >= 'a' && c <= 'z' {
		return true
	}
	if c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}

func TransferChar(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}
