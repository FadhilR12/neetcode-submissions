func productExceptSelf(nums []int) []int {
	var ans []int

	for i, _ := range nums {
		var p int = 1
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}
			p *= nums[j]
		}
		ans = append(ans, p)
	}
	return ans
}