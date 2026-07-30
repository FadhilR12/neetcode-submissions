func productExceptSelf(nums []int) []int {
	var ans []int
	res := 1
	pos := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
			pos = i
		} else if count > 1 {
			res = 0
		} else {
			res *= nums[i]
		}
	}

	for i, v := range nums {
		if count > 0 {
			fmt.Print(count)
			if i != pos || count > 1 {
				ans = append(ans, 0)
				continue
			}
			ans = append(ans, res)
		} else {
			ans = append(ans, (res / v))
		}
	}
	return ans

}