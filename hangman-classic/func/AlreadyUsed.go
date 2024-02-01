package Func

func AlreadyUsed(Used []string, lettre string) bool {
	result := false
	for _, char := range Used {
		if lettre == string(char) {
			result = true
		}
	}
	return result
}
