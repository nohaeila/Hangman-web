// Package Func fournit une collection de fonctions utilitaires.
package Func

import (
	"bufio"
	"os"
)

// State prend un nombre d'essais 'try' en paramètre et retourne une représentation graphique de l'état actuel du pendu.
func State(try int) string {
	// Ouvre le fichier contenant la représentation graphique du pendu.
	file, _ := os.Open("./hangman-classic/hangman.txt")
	
	// Initialise un scanner pour lire le fichier ligne par ligne.
	scanner := bufio.NewScanner(file)
	
	// Calcule les lignes de début et de fin à extraire en fonction du nombre d'essais.
	startLine := (try-1)*7 + 1*try
	endLine := try*7 + try - 1
	
	// Initialise un compteur 'i'.
	i := 1
	
	// Initialise une chaîne de caractères pour stocker le résultat.
	result := ""
	
	// Utilise une boucle pour parcourir chaque ligne du fichier.
	for scanner.Scan() {
		// Si la ligne actuelle est dans la plage définie par startLine et endLine, ajoute la ligne au résultat.
		if i >= startLine && i <= endLine {
			result += scanner.Text() + "\n"
		}
		
		// Incrémente le compteur.
		i++
	}
	
	// Retourne la représentation graphique du pendu basée sur le nombre d'essais.
	return result
}
