package __maxArea

func maxArea(height []int) int {
	curArea, maxArea := 0, 0
	l, r := 0, len(height)-1
	for l < r {
		lh, rh := height[l], height[r]
		curArea = getMin(lh, rh) * (r - l)
		maxArea = getMax(maxArea, curArea)
		if lh < rh {
			l++
		} else {
			r--
		}
	}

	return maxArea
}

func getMin(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func getMax(i, j int) int {
	if i > j {
		return i
	}
	return j
}
