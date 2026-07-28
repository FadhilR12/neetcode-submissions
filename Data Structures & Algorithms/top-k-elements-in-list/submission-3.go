func topKFrequent(nums []int, k int) []int {
	numsMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] += 1
	}

	bucketList := make([][]int, len(nums))
	for num, count := range numsMap {
		bucketList[count - 1] = append(bucketList[count - 1], num)
	}

	ret := make([]int, 0, k)
	for i := len(bucketList) - 1; i >= 0; i-- {
		for _, num := range bucketList[i] {
			ret = append(ret, num)
			if len(ret) == k {
				return ret
			}
		}
	}

	return ret
}