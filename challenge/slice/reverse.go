package slice

func Reverse(a string) string {
	aRunes := []rune(a)
	reversed := []rune{}

	for i := len(aRunes) - 1; i >= 0; i-- {
		reversed = append(reversed, aRunes[i])
	}

	return string(reversed)
}
