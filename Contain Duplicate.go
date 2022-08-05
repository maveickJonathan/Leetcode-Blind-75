func containsDuplicate(nums []int) bool {
	numberMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := numberMap[nums[i]]; ok {
			return true
		}
		numberMap[nums[i]] = 1
	}

	return false
}