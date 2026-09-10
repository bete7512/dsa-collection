func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	slices.Sort(nums)
	results := [][]int{}
	for i := 0; i <= len(nums)-3; i = i + 1 {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[right] + nums[left]
			if sum == 0 {
				results = append(results, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return results
}