package Func

import (
	"fmt"
	"math/rand"
	"time"
)

func Hiderr(word string) []string {
	var hider []string
	for i := 0; i < len(word); i++ {
		hider = append(hider, "_")
	}
	hider = Reveal(word, hider)
	fmt.Println(hider)
	return hider
}
func Reveal(word string, hider []string) []string {
	random := len(word)/2 - 1
	i := 0
	for i < random {
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Intn(len(word))
		if hider[randomNumber] == "_" {
			hider[randomNumber] = string(word[randomNumber])
			i++

		}
	}

	return hider

}
