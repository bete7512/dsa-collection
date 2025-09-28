func longestConsecutive(nums []int) int {
	maps := make(map[int]struct{})
	largestConsecutiveCount := 0

	// put all nums into the map (use struct{} for lighter memory)
	for i := 0; i < len(nums); i++ {
		maps[nums[i]] = struct{}{}
	}

	// iterate over map keys instead of nums
	for num := range maps {
		// only start from the beginning of a sequence
		if _, ok := maps[num-1]; ok {
			continue
		}

		countConsecutive := 1
		k := num + 1
		for {
			if _, ok := maps[k]; !ok {
				break
			}
			countConsecutive++
			k++
		}

		if countConsecutive > largestConsecutiveCount {
			largestConsecutiveCount = countConsecutive
		}
	}

	return largestConsecutiveCount
}
