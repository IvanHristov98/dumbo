package slice

func HasUniqueChars(str string) bool {
	buckets := map[rune]bool{}

	for _, char := range str {
		if _, ok := buckets[char]; ok == true {
			return false
		}

		buckets[char] = true
	}

	return true
}
