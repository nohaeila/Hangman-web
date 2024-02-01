package Func

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func Random(Diff string) string {
	rand.Seed(time.Now().UnixNano())
	t := ""
	s := ""
	var a int
	if Diff == "EASY" {
		a = rand.Intn(37) + 1
		t = "./hangman-classic/words.txt"
	} else if Diff == "NORMAL" {
		a = rand.Intn(23) + 1
		t = "./hangman-classic/words2.txt"
	} else if Diff == "HARD" {
		a = rand.Intn(24) + 1
		t = "./hangman-classic/words3.txt"
	}
	file, err := os.Open(t)
	if err != nil {
		fmt.Println(err)
	}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i++
		if i == a {
			s = scanner.Text()
		}
	}
	println(s)
	return s
}
