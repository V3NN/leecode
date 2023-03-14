package __plusOne

func plusOne(digits []int) []int {
	l := len(digits) - 1
	for i := l; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}

	}

	return append([]int{1}, digits...)
}
