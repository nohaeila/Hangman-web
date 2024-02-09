package hangmanclassic

import (
	"Web/hangman-classic/func"
	"fmt"
)

func Exit() {
	return
}

func InputLetter(word string, try int, used []string, hider []string) {

	Func.State(try)

	var letter string
	fmt.Println("Already used: ", used)
	fmt.Println("\nChoose a letter :")
	fmt.Scanln(&letter)

	letter = Func.ToUpper(letter)
	//input

	if Func.IsAlpha(letter) {
		println("please input a letter")
		InputLetter(word, try, used, hider)
	}
	for _, char := range used {
		if letter == string(char) {
			fmt.Println("Try some thing else , you already use it  !")
			InputLetter(word, try, used, hider)
		}
	}
	used = append(used, letter)
	if len(letter) > 1 {
		if letter == word {
			fmt.Println("Congratulation , you won !")
			Exit()
		} else {
			fmt.Printf("The word %v is incorrect \n", letter)
			if try < 9 {
				try = try + 2
				fmt.Printf("You have %v tries reamaing \n", 10-try)
				InputLetter(word, try, used, hider)
			} else {
				fmt.Println("Not enough tries ")
				fmt.Printf("The correct word was  %v\n", word)
				fmt.Println("Game Over")
				Exit()
			}
		}
	}
	if len(letter) < 1 {
		fmt.Println("Please enter a letter.")
		InputLetter(word, try, used, hider)
	} else {
		inWord := false
		for i := 0; i < len(word); i++ {
			if letter == string(word[i]) {
				inWord = true
			}
		}
		if inWord {
			fmt.Printf("The letter %v is in the word\n", letter)
			for index, char := range word {
				if letter == string(char) {
					hider[index] = letter
				}
			}
		} else {
			fmt.Printf("The letter %v isn't on the word \n", letter)
			if try < 9 {
				try++
				fmt.Printf("You have %v tries remaining  \n", 10-try)
			} else {
				Func.State(10)
				fmt.Println("Not enough tries ")
				fmt.Printf("The correct word was %v\n", word)
				fmt.Println("Game Over")
				Exit()
			}
		}
		fmt.Println(hider)
	}
	stillHider := false
	for index, _ := range word {
		if hider[index] == "_" {
			stillHider = true
			InputLetter(word, try, used, hider)
			break
		}
	}
	if stillHider == false {
		fmt.Println("Congratulation , you won  !")
		Exit()
	}
}
func ReturnHide(word string) []string {
	hider := Func.Hiderr(word)
	fmt.Println(hider)
	return hider
}
func Hangman(Difficulty string) []string {
	var used []string
	tries := 0
	wordLen := Func.Random(Difficulty)
	wordLen = Func.ToUpper(wordLen)
	var hider []string = Func.Hiderr(wordLen)
	InputLetter(wordLen, tries, used, hider)
	return hider
}
