func plusOne(digits []int) []int {
	l := len(digits)
	for i := l - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
        } else {
			digits[i] = 0
			if i == 0 {
				return append([]int{1}, digits...)
			}
        }
	}
    
    return digits
}