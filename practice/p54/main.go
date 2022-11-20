package p54

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	total := m * n
	up_bound, down_bound, left_bound, right_bound := 0, m-1, 0, n-1
	res := make([]int, 0, total)
	for len(res) < total {
		//=>
		if up_bound <= down_bound {
			for j := left_bound; j <= right_bound; j++ {
				res = append(res, matrix[up_bound][j])
			}
			up_bound++
		}

		//down
		if left_bound <= right_bound {
			for j := up_bound; j <= down_bound; j++ {
				res = append(res, matrix[j][right_bound])
			}
			right_bound--
		}

		//<=
		if up_bound <= down_bound {
			for j := right_bound; j >= left_bound; j-- {
				res = append(res, matrix[down_bound][j])
			}
			down_bound--
		}
		//up
		if left_bound <= right_bound {
			for j := down_bound; j >= up_bound; j-- {
				res = append(res, matrix[j][left_bound])
			}
			left_bound++
		}
	}
	return res
}
