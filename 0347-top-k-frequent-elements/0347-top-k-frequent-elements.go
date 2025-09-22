func topKFrequent(nums []int, k int) []int {
	groups := make(map[int]int)
	var groupArrays [][]int
	var results []int
	for _, num := range nums {
		groups[num]++
	}
	for k, v := range groups {
		groupArrays = append(groupArrays, []int{k, v})
	}
	sort.Slice(groupArrays, func(i, j int) bool {
		return groupArrays[i][1] > groupArrays[j][1]
	})
	jsoned, _ := json.Marshal(groupArrays)
	fmt.Println(string(jsoned))
	for i := 0; i < k; i++ {
		results = append(results, groupArrays[i][0])
	}

	return results
}