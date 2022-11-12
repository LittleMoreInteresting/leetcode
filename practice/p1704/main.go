package p1704

func halvesAreAlike(s string) bool {
	n := len(s)
	i, j := 0, n/2
	left, right := 0, 0
	for j < n {
		if isVowel(s[i]) {
			left++
		}
		if isVowel(s[j]) {
			right++
		}
		i++
		j++
	}
	return right == left
}

var vowel = map[byte]bool{
	'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
}

func isVowel(b byte) bool {
	_, ok := vowel[b]
	return ok
}
