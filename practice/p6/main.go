package p6

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	res := make([]string, numRows)
	i, flag := 0, -1
	for _, c := range s {
		res[i] = res[i] + string(c)
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}
	newstr := ""
	for _, str := range res {
		newstr += str
	}
	return newstr
}
