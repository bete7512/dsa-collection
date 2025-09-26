func productExceptSelf(nums []int) []int {
	results := make([]int, len(nums))

	prefixes := make([]int, len(nums))
	postfixes := make([]int, len(nums))
	for i := 0; i < len(nums);{
		if i == 0 {
			prefixes[i] = 1
		}
		if i > 0 {
			prefixes[i] = prefixes[i - 1] * nums[i - 1]
		}
		i++

	}
	for i := 0; i<len(nums);{
		j := len(nums) - 1 - i
		if j == len(nums) - 1 {
			postfixes[j] = 1
		}
		if j < len(nums) - 1 {
			postfixes[j] = postfixes[j + 1] * nums[j + 1]
		}
		i++
	}
	for i := 0; i < len(nums); {
		// prefix := 1
		// postfix := 1
		// for j := i + 1; j < len(nums); {
		// 	postfix *= nums[j]
		// 	j++
		// }

		// for k := 0; k < i; {
		// 	prefix *= nums[k]
		// 	k++
		// }

		results[i] = prefixes[i] * postfixes[i]
		i++
	}

	return results
}