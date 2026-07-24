func hasDuplicate(nums []int) bool {
	var mem map[int]bool = map[int]bool{}
	for _,v := range nums{
		if mem[v]{
			return true
		} else{
			mem[v] = true
		}
	}
	return false
}
