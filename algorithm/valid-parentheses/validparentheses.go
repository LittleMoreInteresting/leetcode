package valid_parentheses

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isValid(s string) bool {
	stringLen := len(s)

	bkts := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < stringLen; i++ {
		if bkt, ok := bkts[s[i]]; ok {
			//右括号则 出栈判断
			stackLen := len(stack)
			if stackLen == 0 || bkt != stack[stackLen-1] {
				return false
			}
			stack = stack[0 : stackLen-1]
		}
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			// 左括号入栈
			stack = append(stack, s[i])
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
