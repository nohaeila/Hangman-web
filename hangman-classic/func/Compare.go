package Func

func Compare(WordToGuess string, Tries int, HiddenWord []string, lettre string) (string, int, string) {
	result := ""

	if len(lettre) > 1 {
		if lettre == WordToGuess {
			result = "Win"
		} else {
			if Tries < 9 {
				Tries = Tries + 2
			} else {
				result = "loose"
			}
		}
	} else {

		inWord := false
		for i := 0; i < len(WordToGuess); i++ {
			if lettre == string(WordToGuess[i]) {
				inWord = true
			}
		}
		if inWord {
			for index, char := range WordToGuess {
				if lettre == string(char) {
					HiddenWord[index] = lettre
				}
			}
		} else {
			if Tries < 9 {
				Tries++
			} else {
				result = "loose"
			}
		}
	}

	stillhider := "false"
	for index, _ := range WordToGuess {
		if HiddenWord[index] == "_" {
			stillhider = "true"
		}
	}
	if stillhider == "false" {
		result = "Win"
	}

	return WordToGuess, Tries, result

}
