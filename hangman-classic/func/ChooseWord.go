package Func

func chooseWord(choosed string) string {
	choosed = ""
	file := ""
	if choosed == "EASY" {
		file = "words.txt"
	}
	if choosed == "NORMAL" {
		file = "words2.txt"
	}
	if choosed == "HARD" {
		file = "words3.txt"
	}
	return file
}
