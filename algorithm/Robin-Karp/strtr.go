package Robin_Karp

import "hash/fnv"

func SearchStr(s, target string) int {
	targetCode := fnv64(target)
	tLen := len(target)
	var i int
	for i = 0; i <= len(s)-tLen; i++ {
		ch := s[i : i+tLen]
		chCode := fnv64(ch)
		if chCode != targetCode {
			continue
		}
		if comStr(ch, target) {
			return i
		}
	}
	return -1
}

func fnv64(key string) uint64 {
	hash := fnv.New64()
	_, _ = hash.Write([]byte(key))
	return hash.Sum64()
}

func comStr(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
