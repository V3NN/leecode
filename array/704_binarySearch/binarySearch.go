package __binarySearch

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (r-l)/2 + l
		mid_val := nums[mid]
		if target < mid_val {
			r = mid - 1
		} else if target > mid_val {
			l = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
