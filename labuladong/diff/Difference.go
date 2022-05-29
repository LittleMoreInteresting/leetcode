package diff

type Difference struct {
	numbs []int
	diff  []int
}

func New(numbs []int) Difference {
	diff := make([]int, len(numbs))
	diff[0] = numbs[0]
	for i := 1; i < len(numbs); i++ {
		diff[i] = numbs[i] - numbs[i-1]
	}
	return Difference{
		numbs: numbs,
		diff:  diff,
	}
}

func (d *Difference) Increment(i, j, inc int) {
	d.diff[i] += inc
	if j+1 < len(d.diff) {
		d.diff[j+1] -= inc
	}
}

func (d *Difference) Result() []int {
	result := make([]int, len(d.diff))
	result[0] = d.diff[0]
	for i := 1; i < len(d.diff); i++ {
		result[i] = result[i-1] + d.diff[i]
	}
	return result
}
