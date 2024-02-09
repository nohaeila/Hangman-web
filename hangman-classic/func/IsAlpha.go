package Func

func IsAlpha(letter string) bool {
	for _, v := range letter {
		if (v < 'A') || (v > 'Z') &&
			(v < 'a') || (v > 'z') {
			return true
		}
	}
	return false
}
