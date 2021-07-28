package stackqueue

import "sync"

type ArrayStack struct {
	array []string
	size  int
	lock  sync.Mutex
}

func (stack *ArrayStack) Push(s string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.array = append(stack.array, s)
	stack.size = stack.size + 1
}

func (stack *ArrayStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if stack.size == 0 {
		panic("stack empty")
	}
	v := stack.array[stack.size-1]
	stack.array = stack.array[0 : stack.size-1]

	/**
		newArray := make([]string, stack.size-1, stack.size-1)
		for i := 0; i < stack.size-1; i++ {
			newArray[i] = stack.array[i]
		}
		stack.array = newArray
	如果切片偏移量向前移动 stack.array[0 : stack.size-1]，
	表明最后的元素已经不属于该数组了，数组变相的缩容了。
	此时，切片被缩容的部分并不会被回收，仍然占用着空间，所以空间复杂度较高，但操作的时间复杂度为：O(1)。
	如果我们创建新的数组 newArray，然后把老数组的元素复制到新数组，就不会占用多余的空间，但移动次数过多，时间复杂度为：O(n)。
	*/
	stack.size = stack.size - 1
	return v
}

func (stack *ArrayStack) Peek() string {
	if stack.size == 0 {
		panic("stack empty")
	}
	v := stack.array[stack.size-1]
	return v
}

func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}
