func topKFrequent(nums []int, k int) []int {
	numsMap := map[int]int{}
	var ans []int

	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] += 1
	}

	for key, _ := range numsMap {
		ans = append(ans, key)
	}

	sort.Slice(ans, func(i, j int) bool { return numsMap[ans[i]] > numsMap[ans[j]] })

	if len(ans) > k {
		ans = ans[:k]
	}
	return ans
}