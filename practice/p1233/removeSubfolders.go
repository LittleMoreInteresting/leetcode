package p1233

import (
	"sort"
)

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	res := []string{}
	for _, f := range folder {
		isC := false
		for _, p := range res {
			if isChild(f, p) {
				isC = true
				break
			}
		}
		if !isC {
			res = append(res, f)
		}
	}

	return res
}

func isChild(f, p string) bool {
	fl, pl := len(f), len(p)
	if fl <= pl {
		return false
	}
	for i := 0; i < pl; i++ {
		if f[i] != p[i] {
			return false
		}
	}
	return f[pl] == '/'
}

/*func removeSubfolders(folder []string) (ans []string) {
	sort.Strings(folder)
	ans = append(ans, folder[0])
	for _, f := range folder[1:] {
		last := ans[len(ans)-1]
		if !strings.HasPrefix(f, last) || f[len(last)] != '/' {
			ans = append(ans, f)
		}
	}
	return
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/remove-sub-folders-from-the-filesystem/solutions/2097563/shan-chu-zi-wen-jian-jia-by-leetcode-sol-0x8d/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。*/
