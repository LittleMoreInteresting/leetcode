package main

import "fmt"

func main() {
	fmt.Println(parseBoolExpr("|(f,t)"))
	fmt.Println(parseBoolExpr("&(t,f)"))
	fmt.Println(parseBoolExpr("|(f,f,f,t)"))
}

func parseBoolExpr(expression string) bool {
	charStack, numStack := []byte{}, []byte{}

	mp := map[byte]byte{
		')': '(',
		'}': '{',
	}
	numBool := map[byte]bool{
		't': true,
		'f': false,
	}
	for i := range expression {
		c := expression[i]
		if c == '&' || c == '|' || c == '!' {
			charStack = append(charStack, c)
			continue
		}
		if left, ok := mp[c]; ok {
			num := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			var newNum bool
			char := charStack[len(charStack)-1]
			charStack = charStack[:len(charStack)-1]
			charNum := []bool{}
			for left != num {
				charNum = append(charNum, numBool[num])
				num = numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
			}
			switch char {
			case '&':
				newNum = charNum[0]
				for _, n := range charNum[1:] {
					newNum = newNum && n
				}
			case '|':
				newNum = charNum[0]
				for _, n := range charNum[1:] {
					newNum = newNum || n
				}
			case '!':
				newNum = !charNum[0]
			}
			if newNum {
				c = 't'
			} else {
				c = 'f'
			}
		}
		if c != ',' {
			numStack = append(numStack, c)
		}

	}
	fmt.Println(string(charStack))
	fmt.Println(string(numStack))
	return numBool[numStack[0]]

}
