package main

import (
	"bufio"
	"fmt"
	"os"
)

// 26
func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := input.Text()
	input_copy := make([]byte, len(str)) // 负责存字母位置和非英文字符
	copy(input_copy, []byte(str))
	output := make([]byte, 0)
	for i := 0; i < 26; i++ {
		for j, value := range input_copy {
			if int(value) == 'a'+i || int(value) == 'A'+i {
				output = append(output, value)
				input_copy[j] = 'a'
			}
		}
	}
	i := 0
	for k, v := range input_copy {
		if v == 'a' {
			input_copy[k] = output[i]
			i++
		}
	}
	fmt.Println(string(input_copy))
}
