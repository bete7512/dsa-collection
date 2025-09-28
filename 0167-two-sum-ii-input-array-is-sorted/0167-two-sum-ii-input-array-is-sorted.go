func twoSum(numbers []int, target int) []int {
	j := len(numbers) - 1
	i := 0
	results := []int{}
	for i < len(numbers) && j > 0 {
		if numbers[i]+numbers[j] > target {
			j--
			continue
		}
		if numbers[i]+numbers[j] < target {
			i++
			continue
		}
		if numbers[i]+numbers[j] == target {
			results = append(results, i+1, j+1)
			break
		}

	}
	return results
}