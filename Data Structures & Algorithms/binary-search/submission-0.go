func search(nums []int, target int) int {
	mid := len(nums) / 2
	left := 0
	right := len(nums) - 1

	for i := 0; i < len(nums); i++ {
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
		mid = (left + right) / 2
	}
	return -1
}