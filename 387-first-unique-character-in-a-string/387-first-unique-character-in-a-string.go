func firstUniqChar(s string) int {
	m := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		c := s[i]
		if _, ok := m[c]; ok {
			m[c] = -1
		} else {
			m[c] = i
		}
	}

	result := 1000000
	for _, v := range m {
		if v != -1 {
			if v < result {
				result = v
			}
		}
	}

	if result == 1000000 {
		return -1
	} else {
		return result
	}

}