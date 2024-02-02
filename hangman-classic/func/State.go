package Func

import (
	"bufio"
	"os"
)

func State(try int) string {
	file, _ := os.Open("./hangman-classic/hangman.txt")
	scanner := bufio.NewScanner(file)
	startLine := (try-1)*7 + 1*try
	endLine := try*7 + try - 1
	i := 1
	result := ""
	for scanner.Scan() {
		if i >= startLine && i <= endLine {
			result += scanner.Text() + "\n"
		}
		i++
	}
	return result
}
