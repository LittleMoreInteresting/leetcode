package main

import "fmt"

//""aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga""
func validPalindrome(s string) bool {

	return validPalindromeRemove(s, false)
}

func validPalindromeRemove(s string, remove bool) bool {
	left, right := 0, len(s)-1
	for left <= right {
		if s[left] == s[right] {
			right--
			left++
			continue
		}
		if !remove {
			remove = true
			fmt.Println("rm1", left, string(s[left]), right, string(s[right]), s[left+1:right+1])
			fmt.Println("rm2", left, string(s[left]), right, string(s[right]), s[left:right])
			if validPalindromeRemove(s[left+1:right+1], remove) ||
				validPalindromeRemove(s[left:right], remove) {
				return true
			}
		}
		return false
	}
	return true
}

func main() {
	v := validPalindrome("aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga")
	fmt.Println(v)
}
