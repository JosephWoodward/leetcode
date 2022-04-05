func lengthOfLastWord(s string) int {
    	l := len(s)

	var result string
	var char string
	for i := l - 1; i >= 0; i-- {
		char = string(s[i])
		if char == " " {
			if result != "" {
				return len(result)
			}
			continue
		}
		if char != " " {
			result += char
		}
	}

	return len(result)

}