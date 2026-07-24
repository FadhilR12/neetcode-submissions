func twoSum(nums []int, target int) []int {
    var ans []int
	numMap := map[int]int{}

	for i, val := range nums {
		diff := target - val
		mapVal, ok := numMap[diff]
		if ok {
			ans = append(ans, mapVal, i)
			return ans
		}
		numMap[val] = i
	}
	return ans
}
