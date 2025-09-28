func longestConsecutive(nums []int) int {
	maps := make(map[int]int)
	largestConsecutiveCount := 0
	for i := 0; i < len(nums); i++ {
		maps[nums[i]] = i
	}

	for num := range maps{
		if _, ok := maps[num-1]; ok {
			continue
		}

		countConsecutive := 1
		k := num + 1
		for {
			_, ok := maps[k]
			if !ok {
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