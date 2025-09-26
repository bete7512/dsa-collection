func productExceptSelf(nums []int) []int {
	results := make([]int, len(nums))

	prefixes := make([]int, len(nums))
	postfixes := make([]int, len(nums))
	for i := 0; i < len(nums); {
		if i == 0 {
			prefixes[i] = 1
		}
		if i > 0 {
			prefixes[i] = prefixes[i-1] * nums[i-1]
		}

		j := len(nums) - 1 - i
		if j == len(nums)-1 {
			postfixes[j] = 1
		}
		if j < len(nums)-1 {
			postfixes[j] = postfixes[j+1] * nums[j+1]
		}
		i++

	}
	for i := 0; i < len(nums); {

		results[i] = prefixes[i] * postfixes[i]
		i++
	}

	return results
}