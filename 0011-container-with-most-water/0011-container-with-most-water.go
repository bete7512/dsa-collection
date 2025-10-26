func maxArea(height []int) int {
	maxarea := 0
	j := len(height) - 1
	i := 0
	for i < len(height) && j > i {
		area := (j - i) * (min(height[i], height[j]))
		if area > maxarea {
			maxarea = area
		}

		if height[i] < height[j] {
			i++
		} else if height[j] < height[i] {
			j--
		} else {
			j--
		}
	}
	return maxarea
}