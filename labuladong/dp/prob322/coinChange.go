package prob322

func coinChange(coins []int, amount int) int {
	mem := make([]int, amount+1)
	mem[0] = 0
	for i := 1; i < len(mem); i++ {
		mem[i] = amount + 1
	}
	for i := 0; i < len(mem); i++ {
		for _, coin := range coins {
			if i < coin {
				continue
			}
			mem[i] = min(mem[i], 1+mem[i-coin])
		}
	}
	if mem[amount] == amount+1 {
		return -1
	}
	return mem[amount]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
