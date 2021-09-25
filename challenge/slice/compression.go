package slice

import "fmt"

func Compress(a string) string {
	compressed := ""
	repetitions := 0
	prev := rune(0)

	for _, r := range a {
		if r == prev {
			repetitions++
			continue
		}

		compressed += compressRune(prev, repetitions)
		repetitions = 1
		prev = r
	}

	compressed += compressRune(prev, repetitions)

	if len(compressed) == len(a) {
		return a
	}

	return compressed
}

func compressRune(r rune, repetitions int) string {
	if repetitions <= 0 {
		return ""
	}

	return fmt.Sprintf("%s%d", string(r), repetitions)
}
