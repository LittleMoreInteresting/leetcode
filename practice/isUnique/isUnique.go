package isUnique

func isUnique(astr string) bool {
	mark := 0
	for _, c := range astr {
		bit := c - 'a'
		if mark&(1<<bit) != 0 {
			return false
		} else {
			mark |= (1 << bit)
		}
	}
	return true
}
