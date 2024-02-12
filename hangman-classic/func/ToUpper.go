// Package Func fournit une collection de fonctions utilitaires.
package Func
// ToUpper prend une lettre en paramètre et retourne la version en majuscules de cette lettre.
func ToUpper(letter string) string {
	// Convertit la lettre en une séquence de runes pour permettre la modification individuelle.
	a := []rune(letter)
	
	// Utilise une boucle pour parcourir chaque rune de la lettre.
	for i := 0; i < len(a); i++ {
		// Vérifie si la rune actuelle est une minuscule.
		if a[i] >= 'a' && a[i] <= 'z' {
			// Si vrai, convertit la minuscule en majuscule en ajustant la valeur de la rune.
			a[i] -= 32
		}
	}
	
	// Retourne la lettre modifiée en majuscules.
	return string(a)
}
