package Func

func ToUpper(letter string) string {
	a := []rune(letter)
	for i := 0; i < len(a); i++ {
		if a[i] >= 'a' && a[i] <= 'z' {
			a[i] -= 32
		}
	}
	return string(a)
}
