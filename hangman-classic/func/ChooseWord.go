// Package Func fournit une collection de fonctions utilitaires.
package Func
// chooseWord prend un niveau de difficulté 'choosed' en paramètre et retourne le chemin du fichier de mots correspondant au niveau de difficulté spécifié.
func chooseWord(choosed string) string {
	// Réinitialise la variable 'choosed'.
	choosed = ""
	
	// Initialise la variable 'file'.
	file := ""
	
	// Utilise une série d'instructions if pour déterminer le fichier de mots en fonction du niveau de difficulté.
	if choosed == "EASY" {
		file = "words.txt"
	}
	if choosed == "NORMAL" {
		file = "words2.txt"
	}
	if choosed == "HARD" {
		file = "words3.txt"
	}
	
	// Retourne le chemin du fichier de mots correspondant au niveau de difficulté spécifié.
	return file
}

