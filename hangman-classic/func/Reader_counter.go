// Package Func fournit une collection de fonctions utilitaires.
package Func

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Random prend un niveau de difficulté 'Diff' en paramètre et retourne un mot aléatoire correspondant au niveau de difficulté spécifié.
func Random(Diff string) string {
	// Initialise le générateur de nombres aléatoires avec une nouvelle graine basée sur le temps actuel.
	rand.Seed(time.Now().UnixNano())
	
	// Initialise les variables 't' et 's' pour le chemin du fichier de mots et le mot sélectionné, respectivement.
	t := ""
	s := ""
	
	// Initialise 'a' pour stocker un nombre aléatoire qui sera utilisé pour sélectionner un mot du fichier.
	var a int
	
	// Utilise une instruction if-else pour déterminer le niveau de difficulté et choisir le fichier de mots approprié.
	if Diff == "FACILE" {
		a = rand.Intn(37) + 1
		t = "./hangman-classic/words.txt"
	} else if Diff == "NORMAL" {
		a = rand.Intn(23) + 1
		t = "./hangman-classic/words2.txt"
	} else if Diff == "DIFFICILE" {
		a = rand.Intn(24) + 1
		t = "./hangman-classic/words3.txt"
	}
	
	// Ouvre le fichier correspondant au niveau de difficulté.
	file, err := os.Open(t)
	if err != nil {
		fmt.Println(err)
	}
	
	// Initialise un compteur 'i'.
	i := 0
	
	// Utilise un scanner pour lire chaque ligne du fichier.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i++
		// Si le compteur atteint le nombre aléatoire 'a', sélectionne le mot correspondant.
		if i == a {
			s = scanner.Text()
		}
	}
	
	// Affiche le mot sélectionné à des fins de débogage.
	println(s)
	
	// Retourne le mot sélectionné.
	return s
}
