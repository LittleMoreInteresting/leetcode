package main

import (
	"bufio"
	"fmt"
	"os"
)

// 生成字母表
func getTable(key []byte, tableMap map[byte]byte) {
	tmpMap := make(map[byte]int)
	k := 0
	for _, ch := range key {
		if _, ok := tmpMap[ch]; !ok {
			tableMap[byte('a'+k)] = ch
			tmpMap[ch] = 1
			k++
		}
	}
	tmpMap2 := make(map[int]byte)
	k = 0
	for i := 0; i < 26; i++ {
		if _, ok := tmpMap[byte('a'+i)]; !ok {
			tmpMap2[k] = byte('a' + i)
			k++
		}
	}
	k = 0
	for i := 'a' + len(tmpMap); i <= 'z'; i++ {
		tableMap[byte(i)] = tmpMap2[k]
		k++
	}
}

func getRes(chs []byte, tableMap map[byte]byte) {
	for _, ch := range chs {
		fmt.Print(string(tableMap[ch]))
	}
	fmt.Print("\n")
}

// 输入单词，去除重复的字母，按顺序输出，26个字母没出现的再接着输出，对于a~z生成字母转换表
// 多组数据，输入仅含小写
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	i := 0
	tableMap := make(map[byte]byte)
	for scanner.Scan() {
		i++
		str := scanner.Text()
		if i%2 == 1 { // 密钥
			getTable([]byte(str), tableMap)
		} else {
			getRes([]byte(str), tableMap)
		}
	}
}
