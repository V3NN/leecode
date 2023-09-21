package __firstBadVersion

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 *
 */

func firstBadVersion(n int) int {
	l := 0

	for l <= n {
		mid := (l + n) / 2
		if isBadVersion(mid) {
			n = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func isBadVersion(version int) bool {
	if version == 5 {
		return true
	}
	return false
}
