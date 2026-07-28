func groupAnagrams(strs []string) [][]string {
	ans := map[string][]string{}
	var value [][]string
	for _, v := range strs {
		char := strings.Split(v, "")
		sort.Slice(char, func(i, j int) bool { return char[i] < char[j] })
		k := strings.Join(char, "")
		ans[k] = append(ans[k], v)
	}
	for _, val := range ans {
		value = append(value, val)
	}
	return value
}