package main

import "sort"

func main() {

}
func customSortString(order string, s string) string {
	sortIdx := map[byte]int{}
	for i, c := range order {
		sortIdx[byte(c)] = i + 1
	}
	str := []byte(s)
	sort.Slice(str, func(i, j int) bool { return sortIdx[str[i]] < sortIdx[str[j]] })
	return string(str)
}

func customSortString1(order, s string) string {
	val := map[byte]int{}
	for i, c := range order {
		val[byte(c)] = i + 1
	}
	t := []byte(s)
	sort.Slice(t, func(i, j int) bool { return val[t[i]] < val[t[j]] })
	return string(t)
}

/*作者：力扣官方题解
链接：https://leetcode.cn/problems/custom-sort-string/solutions/1963410/zi-ding-yi-zi-fu-chuan-pai-xu-by-leetcod-1qvf/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。*/
