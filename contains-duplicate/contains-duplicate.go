func containsDuplicate(nums []int) bool {
    m := make(map[int]struct{})
	for _, num := range nums {
		_, exists := m[num]
		if exists {
			return true
		}

        m[num] = struct{}{}
	}

	return false
}