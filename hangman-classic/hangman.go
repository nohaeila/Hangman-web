package hangmanclassic

import (
	"Web/hangman-classic/func"
	"fmt"
)

// Exit termine le programme.
func Exit() {
	return
}

// InputLetter prend en paramètres le mot à deviner 'word', le nombre d'essais 'try', la liste des lettres déjà utilisées 'used', et la liste de lettres cachées 'hider'.
// La fonction gère la logique de jeu en prenant l'entrée de l'utilisateur pour deviner une lettre et met à jour l'état du jeu en conséquence.
func InputLetter(word string, try int, used []string, hider []string) {

	// Affiche l'état actuel du pendu en fonction du nombre d'essais.
	Func.State(try)

	var letter string
	fmt.Println("Already used: ", used)
	fmt.Println("\nChoose a letter :")
	fmt.Scanln(&letter)

	// Convertit la lettre en majuscules.
	letter = Func.ToUpper(letter)

	// Vérifie si la lettre est alphabétique.
	if Func.IsAlpha(letter) {
		println("please input a letter")
		InputLetter(word, try, used, hider)
	}

	// Vérifie si la lettre a déjà été utilisée.
	for _, char := range used {
		if letter == string(char) {
			fmt.Println("Try something else, you already used it!")
			InputLetter(word, try, used, hider)
		}
	}

	// Ajoute la lettre à la liste des lettres déjà utilisées.
	used = append(used, letter)

	// Gère les cas où la lettre est trop longue, est correcte ou incorrecte.
	if len(letter) > 1 {
		if letter == word {
			fmt.Println("Congratulations, you won!")
			Exit()
		} else {
			fmt.Printf("The word %v is incorrect \n", letter)
			if try < 9 {
				try = try + 2
				fmt.Printf("You have %v tries remaining \n", 10-try)
				InputLetter(word, try, used, hider)
			} else {
				fmt.Println("Not enough tries ")
				fmt.Printf("The correct word was %v\n", word)
				fmt.Println("Game Over")
				Exit()
			}
		}
	}

	// Gère le cas où la lettre est vide.
	if len(letter) < 1 {
		fmt.Println("Please enter a letter.")
		InputLetter(word, try, used, hider)
	} else {
		inWord := false
		// Vérifie si la lettre est présente dans le mot à deviner.
		for i := 0; i < len(word); i++ {
			if letter == string(word[i]) {
				inWord = true
			}
		}
		if inWord {
			// Si vrai, met à jour les lettres cachées avec la lettre correcte.
			fmt.Printf("The letter %v is in the word\n", letter)
			for index, char := range word {
				if letter == string(char) {
					hider[index] = letter
				}
			}
		} else {
			// Si faux, gère le cas où la lettre est incorrecte.
			fmt.Printf("The letter %v isn't in the word \n", letter)
			if try < 9 {
				try++
				fmt.Printf("You have %v tries remaining  \n", 10-try)
			} else {
				// Si le nombre d'essais est épuisé, affiche le pendu complet.
				Func.State(10)
				fmt.Println("Not enough tries ")
				fmt.Printf("The correct word was %v\n", word)
				fmt.Println("Game Over")
				Exit()
			}
		}
		fmt.Println(hider)
	}

	// Vérifie si des lettres cachées subsistent dans le mot à deviner.
	stillHider := false
	for index, _ := range word {
		if hider[index] == "_" {
			stillHider = true
			InputLetter(word, try, used, hider)
			break
		}
	}
	if !stillHider {
		// Si aucune lettre cachée n'est présente, le joueur a gagné.
		fmt.Println("Congratulations, you won!")
		Exit()
	}
}

// ReturnHide prend un mot à deviner 'word' en paramètre et retourne la liste de lettres cachées correspondante.
func ReturnHide(word string) []string {
	hider := Func.Hiderr(word)
	fmt.Println(hider)
	return hider
}

// Hangman initialise une partie de pendu en fonction du niveau de difficulté spécifié.
// La fonction retourne la liste de lettres cachées.
func Hangman(Difficulty string) []string {
	var used []string
	tries := 0
	wordLen := Func.Random(Difficulty)
	wordLen = Func.ToUpper(wordLen)
	var hider []string = Func.Hiderr(wordLen)
	InputLetter(wordLen, tries, used, hider)
	return hider
}
