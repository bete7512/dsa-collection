func groupAnagrams(strs []string) [][]string {
	groupMaps := make(map[string][]string)
	var results [][]string
    for _, str := range strs {
        s := []rune(str)
        sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
        groupMaps[string(s)] = append(groupMaps[string(s)], str)
    }
	for _, v := range groupMaps {
		results = append(results, v)
	}

	return results
}
