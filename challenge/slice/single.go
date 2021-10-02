package slice

import "fmt"

func FindDiffChar(a, b string) (rune, error) {
	counts := map[rune]int{}

	for _, r := range a {
		setDefaultRuneCount(counts, r)
		counts[r]++
	}

	for _, r := range b {
		setDefaultRuneCount(counts, r)
		counts[r]--
	}

	diffRune := rune(0)
	diffCount := 0

	for r, count := range counts {
		if count == -1 || count == 1 {
			diffRune = r
			diffCount++
		}
	}

	if diffCount == 0 {
		return rune(0), fmt.Errorf("No different runes")
	}

	if diffCount > 1 {
		return rune(0), fmt.Errorf("More than one different runes")
	}

	return diffRune, nil
}

func setDefaultRuneCount(counts map[rune]int, r rune) {
	if _, exists := counts[r]; !exists {
		counts[r] = 0
	}
}
