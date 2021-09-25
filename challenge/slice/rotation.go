package slice

func IsRotation(a, b string) bool {
	doubleA := a + a

	return isSubstr(doubleA, b)
}

func isSubstr(super, sub string) bool {
	j := 0

	for i, b := range []byte(super) {
		if j == len(sub) {
			return true
		}

		if len(super[i:]) < len(sub[j:]) {
			return false
		}

		if b != sub[j] {
			j = 0
		} else {
			j++
		}
	}

	return true
}
