// Package Func fournit une collection de fonctions utilitaires.
package Func

import (
	"fmt"
	"math/rand"
	"time"
)

// Hiderr prend un mot à cacher 'word' en paramètre et retourne une liste de lettres cachées initiale.
func Hiderr(word string) []string {
	// Initialise une liste de lettres cachées 'hider'.
	var hider []string

	// Remplit 'hider' avec des underscores '_' pour chaque lettre du mot 'word'.
	for i := 0; i < len(word); i++ {
		hider = append(hider, "_")
	}

	// Appelle la fonction Reveal pour révéler quelques lettres aléatoires dans 'hider'.
	hider = Reveal(word, hider)

	// Affiche 'hider' à des fins de débogage.
	fmt.Println(hider)

	// Retourne la liste de lettres cachées 'hider'.
	return hider
}

// Reveal prend un mot 'word' et une liste de lettres cachées 'hider' en paramètres.
// Elle révèle un nombre aléatoire de lettres du mot dans la liste et la retourne mise à jour.
func Reveal(word string, hider []string) []string {
	// Calcule un nombre aléatoire qui représente la moitié de la longueur du mot moins 1.
	random := len(word)/2 - 1

	// Initialise un compteur 'i'.
	i := 0

	// Boucle jusqu'à ce que le compteur atteigne la valeur aléatoire.
	for i < random {
		// Initialise le générateur de nombres aléatoires avec une nouvelle graine basée sur le temps actuel.
		rand.Seed(time.Now().UnixNano())

		// Génère un nombre aléatoire entre 0 et la longueur du mot.
		randomNumber := rand.Intn(len(word))

		// Vérifie si la lettre correspondante dans 'hider' est encore cachée.
		if hider[randomNumber] == "_" {
			// Si vrai, révèle la lettre correspondante du mot dans 'hider'.
			hider[randomNumber] = string(word[randomNumber])

			// Incrémente le compteur.
			i++
		}
	}

	// Retourne la liste de lettres cachées 'hider' mise à jour.
	return hider
}
