// Package Func fournit une collection de fonctions utilitaires.
package Func
// AlreadyUsed vérifie si une lettre donnée a déjà été utilisée en la comparant à une liste de lettres utilisées.
func AlreadyUsed(Used []string, lettre string) bool {
	// Initialise la variable résultat à false, en supposant que la lettre n'a pas encore été utilisée.
	result := false
	
	// Itère à travers chaque caractère dans la liste Used.
	for _, char := range Used {
		// Vérifie si le caractère actuel dans la liste est égal à la lettre donnée.
		if lettre == string(char) {
			// Si une correspondance est trouvée, met à jour le résultat à true.
			result = true
		}
	}
	
	// Retourne le résultat final indiquant si la lettre a déjà été utilisée.
	return result
}
