// Package Func fournit une collection de fonctions utilitaires.
package Func
// Compare prend en paramètres un mot à deviner 'WordToGuess', un nombre d'essais 'Tries', une liste de lettres cachées 'HiddenWord', et une lettre 'lettre'.
// La fonction retourne le mot à deviner mis à jour, le nombre d'essais mis à jour, et un résultat ("Win", "loose" ou une chaîne vide).
func Compare(WordToGuess string, Tries int, HiddenWord []string, lettre string) (string, int, string) {
	// Initialise la variable 'result' à une chaîne vide.
	result := ""

	// Vérifie si la longueur de 'lettre' est supérieure à 1, indiquant une tentative de deviner le mot complet.
	if len(lettre) > 1 {
		// Vérifie si 'lettre' est égal au mot à deviner 'WordToGuess'.
		if lettre == WordToGuess {
			// Si vrai, met à jour 'result' avec "Win".
			result = "Win"
		} else {
			// Si faux, vérifie si le nombre d'essais 'Tries' est inférieur à 9.
			if Tries < 9 {
				// Si vrai, incrémente 'Tries' de 2.
				Tries = Tries + 2
			} else {
				// Si faux, met à jour 'result' avec "loose".
				result = "loose"
			}
		}
	} else {
		// Si 'lettre' est une seule lettre.

		// Initialise 'inWord' à false, indiquant que la lettre n'est pas dans le mot à deviner.
		inWord := false

		// Parcourt le mot à deviner 'WordToGuess'.
		for i := 0; i < len(WordToGuess); i++ {
			// Vérifie si 'lettre' est égal à la lettre actuelle dans le mot.
			if lettre == string(WordToGuess[i]) {
				// Si vrai, met à jour 'inWord' à true.
				inWord = true
			}
		}

		// Vérifie si la lettre est dans le mot.
		if inWord {
			// Si vrai, parcourt le mot à deviner pour mettre à jour les lettres cachées.
			for index, char := range WordToGuess {
				if lettre == string(char) {
					HiddenWord[index] = lettre
				}
			}
		} else {
			// Si faux, vérifie si le nombre d'essais 'Tries' est inférieur à 9.
			if Tries < 9 {
				// Si vrai, incrémente 'Tries'.
				Tries++
			} else {
				// Si faux, met à jour 'result' avec "loose".
				result = "loose"
			}
		}
	}

	// Vérifie si des lettres cachées subsistent dans le mot à deviner 'WordToGuess'.
	stillhider := "false"
	for index, _ := range WordToGuess {
		if HiddenWord[index] == "_" {
			stillhider = "true"
		}
	}

	// Si aucune lettre cachée n'est présente, met à jour 'result' avec "Win".
	if stillhider == "false" {
		result = "Win"
	}

	// Retourne le mot à deviner mis à jour, le nombre d'essais mis à jour, et le résultat.
	return WordToGuess, Tries, result
}
