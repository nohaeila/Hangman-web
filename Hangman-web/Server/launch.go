package Server

import (
	"Web/Hangman-classic/func"
	"fmt" // Replace "github.com/username/repo" with the actual package path
	"html/template"
	"net/http"
)

type Hangman struct {
	WordToGuess string
	HiddenWord  []string
	UserInput   string
	Tries       int
	Used        []string
	Jose        string
	Pseudo      string
	Difficulty  string
	Result      string
	Suggest     string
}

var Data Hangman

func Display(w http.ResponseWriter, r *http.Request) {
	custTemplate, err := template.ParseFiles("Hangman-web/templates/html/accueil.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	buttonValue := r.FormValue("bouton")
	username := r.FormValue("nom_utilisateur")
	if buttonValue != "" {
		fmt.Printf("Le Bouton prends une valeur : %s\n", buttonValue)
		Data.WordToGuess = Func.Random(buttonValue)
		Data.WordToGuess = Func.ToUpper(Data.WordToGuess)
		Data.HiddenWord = Func.Hiderr(Data.WordToGuess)
		Data.Tries = 0
		Data.Difficulty = buttonValue
		Data.Used = []string{""}
		Data.Jose = Func.DisplayImage(w, r, Data.Tries)
		Data.Suggest = ""
	}
	if username != "" {
		fmt.Printf("Le pseudo de l'utilisateur est  : %s\n", username)
		Data.Pseudo = username
		http.Redirect(w, r, "/index", http.StatusFound)
	}
	err = custTemplate.Execute(w, Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func DisplayIndex(w http.ResponseWriter, r *http.Request) {
	custTemplate, err := template.ParseFiles("Hangman-web/templates/html/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	lettre := r.FormValue("lettre")
	if lettre != "" {
		Data.Suggest = ""
		if Func.IsAlpha(lettre) == true {
			Data.Suggest = Func.Suggest(1)
		} else {
			lettre = Func.ToUpper(lettre)
			if Func.AlreadyUsed(Data.Used, lettre) == true {
				Data.Suggest = Func.Suggest(2)
			} else {
				Data.Used = append(Data.Used, lettre)
				println("la lettre est entrée est ", lettre)
				if lettre != "" {
					Data.WordToGuess, Data.Tries, Data.Result = Func.Compare(Data.WordToGuess, Data.Tries, Data.HiddenWord, lettre)
					if Data.Result == "Win" {
						http.Redirect(w, r, "/win", http.StatusFound)
					} else if Data.Result == "loose" {
						http.Redirect(w, r, "/loose", http.StatusFound)
					}
				}
			}
		}
		Data.Jose = Func.DisplayImage(w, r, Data.Tries)
	}
	err = custTemplate.Execute(w, Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func DisplayWin(w http.ResponseWriter, r *http.Request) {
	custTemplate, err := template.ParseFiles("Hangman-web/templates/html/win.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = custTemplate.Execute(w, Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func DisplayLoose(w http.ResponseWriter, r *http.Request) {
	custTemplate, err := template.ParseFiles("Hangman-web/templates/html/loose.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = custTemplate.Execute(w, Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func StartServer() {
	println("Server started on http://localhost:8080/")
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./templates/css"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./assets/images"))))
	http.Handle("/musique/", http.StripPrefix("/musique/", http.FileServer(http.Dir("./assets/musique/"))))
	http.HandleFunc("/", Display)
	http.HandleFunc("/index", DisplayIndex)
	http.HandleFunc("/win", DisplayWin)
	http.HandleFunc("/loose", DisplayLoose)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
