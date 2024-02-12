// Package Func fournit une collection de fonctions utilitaires.
package Func
// IsAlpha prend une lettre en paramètre et retourne un booléen indiquant si la lettre est alphabétique.
func IsAlpha(letter string) bool {
	// Itère à travers chaque caractère dans la lettre fournie.
	for _, v := range letter {
		// Vérifie si le caractère n'est pas compris entre 'A' et 'Z' (majuscules) ou entre 'a' et 'z' (minuscules).
		if (v < 'A' || v > 'Z') &&
			(v < 'a' || v > 'z') {
			// Si vrai, la lettre n'est pas alphabétique, retourne true.
			return true
		}
	}
	
	// Si aucune condition n'est satisfaite, la lettre est alphabétique, retourne false.
	return false
}
