func reversePrefix(word string, ch byte) string {
	w := []byte(word)
	index := bytes.IndexByte(w, ch)
	sub := w[:index+1]

	p2 := index
	for i := 0; i < p2; i++ {
		sub[i], sub[p2] = sub[p2], sub[i]
		p2--
	}

	return string(w)
}
