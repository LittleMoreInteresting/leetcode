package p435

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	total, prev := 0, intervals[0][1]
	for i := 1; i < n; i++ {
		if intervals[i][0] < prev {
			total++
		} else {
			prev = intervals[i][1]
		}
	}
	return total
}
