func lengthOfLastWord(s string) int {
	pos := len(s) - 1

	var result string
	var char rune
	for ; pos >= 0; pos-- {
		char = rune(s[pos])
		if unicode.IsSpace(char) {
			if result != "" {
				return len(result)
			}
			continue
		} else {
			result += string(char)
		}
	}

	return len(result)


}