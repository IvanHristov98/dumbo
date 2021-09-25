package slice

func IsPermutation(a string, b string) bool {
	buckets := map[rune]int{}

	extractLetters(a, buckets)
	takeLetters(b, buckets)

	return areBucketsEmpty(buckets)
}

func extractLetters(a string, buckets map[rune]int) {
	for _, r := range a {
		if _, ok := buckets[r]; !ok {
			buckets[r] = 0
		}

		buckets[r]++
	}
}

func takeLetters(a string, buckets map[rune]int) {
	for _, r := range a {
		count, ok := buckets[r]

		if !ok || count <= 0 {
			buckets[r] = 0
		}

		buckets[r]--
	}
}

func areBucketsEmpty(buckets map[rune]int) bool {
	for _, count := range buckets {
		if count != 0 {
			return false
		}
	}

	return true
}
