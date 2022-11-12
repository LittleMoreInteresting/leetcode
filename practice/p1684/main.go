package p1684

var allow map[byte]bool

func countConsistentStrings(allowed string, words []string) int {
	allow = make(map[byte]bool)
	for i := range allowed {
		allow[allowed[i]] = true
	}
	count := 0
	for _, wd := range words {
		if isAllow(wd) {
			count++
		}
	}
	return count
}
func countConsistentStrings1(allowed string, words []string) (res int) {
	mask := 0
	for _, c := range allowed {
		mask |= 1 << (c - 'a')
	}
	for _, word := range words {
		mask1 := 0
		for _, c := range word {
			mask1 |= 1 << (c - 'a')
		}
		if mask1|mask == mask {
			res++
		}
	}
	return
	//作者：力扣官方题解
}

func isAllow(wd string) bool {
	for i := range wd {
		if _, ok := allow[wd[i]]; !ok {
			return false
		}
	}
	return true
}
