// Package Func fournit une collection de fonctions utilitaires.
package Func
// Suggest prend un code d'erreur 'Error' en paramètre et retourne un message de suggestion correspondant.
func Suggest(Error int) string {
	// Initialise la variable 'Message'.
	var Message string
	
	// Utilise une instruction switch pour déterminer le message de suggestion en fonction du code d'erreur.
	switch Error {
	case 1:
		// Si le code d'erreur est 1, le message est "Please enter a letter".
		Message = "Veuillez entrer une lettre"
	case 2:
		// Si le code d'erreur est 2, le message est "Letter already used, Please try something else".
		Message = "Lettre déjà utilisée, veuillez essayer autre chose"
	default:
		// Si le code d'erreur ne correspond à aucune des conditions, retourne une chaîne vide.
		return ""
	}
	
	// Retourne le message de suggestion correspondant au code d'erreur.
	return Message
}
