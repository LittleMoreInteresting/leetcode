package p2342

func maximumSum(nums []int) int {
	max := -1
	res := make(map[int][2]int)
	for i := 0; i < len(nums); i++ {
		sv := getSun(nums[i])
		if _, ok := res[sv]; !ok {
			res[sv] = [2]int{nums[i], -1}
		} else {
			arr := res[sv]
			if nums[i] > arr[0] {
				arr[1] = arr[0]
				arr[0] = nums[i]
			} else if nums[i] > arr[1] {
				arr[1] = nums[i]
			}
			res[sv] = arr
		}
	}
	for _, arr := range res {
		if arr[1] == -1 {
			continue
		}
		sumVal := arr[0] + arr[1]
		if sumVal > max {
			max = sumVal
		}
	}
	return max
}

func getSun(num int) int {
	s := 0
	for num > 0 {
		s += num % 10
		num = num / 10
	}
	s += num
	return s
}
