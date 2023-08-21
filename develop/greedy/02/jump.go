package main

import (
	"fmt"
)

func main() {
	/*scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	initNum := strings.Split(text, " ")
	length, _ := strconv.Atoi(initNum[0])
	stone, _ := strconv.Atoi(initNum[1])
	remove, _ := strconv.Atoi(initNum[2])
	stones := make([]int, stone)
	for i := 0; i < stone; i++ {
		var li int
		fmt.Scan(&li)
		stones[i] = li
	}*/
	length, remove, stone := 25, 2, 5
	stones := []int{2, 11, 14, 17, 21}
	fmt.Println(length, remove, stone)
	fmt.Println(stones)
	left, right := 1, length
	for left <= right {
		mid := left + (right-left)/2
		count := removeStoneCount(stones, length, mid)
		fmt.Println(right, left, mid, count)

		if count <= remove {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Println(right)
}

func removeStoneCount(s []int, length int, gap int) int {
	if len(s) == 0 {
		return 0
	}
	count := 0
	pre := 0
	for _, v := range s {
		if v-pre < gap {
			count += 1
		} else {
			pre = v
		}
	}
	if length-s[len(s)-1] < gap {
		count += 1
	}
	return count
}
