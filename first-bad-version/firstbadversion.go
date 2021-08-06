package first_bad_version

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	l := 0
	r := n
	last := -1
	for {
		mid := (l + r) / 2
		if isBadVersion(mid) {
			last = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
		if l > r {
			return last
		}
	}
}

func isBadVersion(version int) bool {
	if version >= 4 {
		return true
	}
	return false
}
