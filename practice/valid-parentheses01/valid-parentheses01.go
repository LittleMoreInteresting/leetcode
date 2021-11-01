package valid_parentheses01

import "fmt"

func isValid(s string) bool {
	l := len(s)
	rpar := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	mockStack := []byte{}
	for i := 0; i < l; i++ {
		if b, ok := rpar[s[i]]; ok {
			// 右括号 入栈
			mockStack = append(mockStack, b)
			fmt.Printf("mockStack:%v\n", mockStack)
			continue
		}
		if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			//左括号出栈
			stackLen := len(mockStack)
			if stackLen == 0 || mockStack[stackLen-1] != s[i] {
				return false
			}
			mockStack = mockStack[0 : stackLen-1]
			fmt.Printf("mockStack:%v\n", mockStack)
		}

	}
	if len(mockStack) != 0 {
		fmt.Printf("mockStack:%v\n", mockStack)
		return false
	}
	return true

}
