func hasDuplicate(nums []int) bool {
	for i := len(nums) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] == nums[i] {
				return true
			}
		}
	}
	return false
}
