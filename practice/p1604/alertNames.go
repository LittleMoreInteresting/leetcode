package p1604

import (
	"sort"
)

func alertNames(keyName []string, keyTime []string) []string {
	count := map[string][]int{}
	for i, name := range keyName {
		t := keyTime[i]
		hour := int(t[0]-'0')*10 + int(t[1]-'0')
		minute := int(t[3]-'0')*10 + int(t[4]-'0')
		count[name] = append(count[name], hour*60+minute)
	}
	res := []string{}
	for k, v := range count {
		if len(v) >= 3 {
			sort.Ints(v)
			for i, t := range v[2:] {
				if t-v[i] >= 60 {
					res = append(res, k)
				}
			}
		}
	}
	sort.Strings(res)
	return res
}
