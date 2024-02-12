package Server

import (
	"Web/hangman-classic/func"
	"fmt"
	"html/template"
	"net/http"
)

// Hangman structure représente l'état du jeu de pendu.
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

// Data contient l'état actuel du jeu.
var Data Hangman

// Display gère la logique d'affichage de la page d'accueil du jeu de pendu.
func Display(w http.ResponseWriter, r *http.Request) {
	// Parse le fichier HTML du template.
	custTemplate, err := template.ParseFiles("Hangman-web/templates/accueil.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Récupère la valeur du bouton et du nom d'utilisateur depuis le formulaire.
	buttonValue := r.FormValue("bouton")
	username := r.FormValue("nom_utilisateur")

	// Traite la valeur du bouton.
	if buttonValue != "" {
		fmt.Printf("Le Bouton prends une valeur : %s\n", buttonValue)
		// Initialise les données pour une nouvelle partie.
		Data.WordToGuess = Func.Random(buttonValue)
		Data.WordToGuess = Func.ToUpper(Data.WordToGuess)
		Data.HiddenWord = Func.Hiderr(Data.WordToGuess)
		Data.Tries = 0
		Data.Difficulty = buttonValue
		Data.Used = []string{""}
		Data.Jose = Func.DisplayImage(w, r, Data.Tries)
		Data.Suggest = ""
	}

	// Traite le nom d'utilisateur.
	if username != "" {
		fmt.Printf("Le pseudo de l'utilisateur est  : %s\n", username)
		Data.Pseudo = username
		// Redirige vers la page principale du jeu.
		http.Redirect(w, r, "/index", http.StatusFound)
	}

	// Exécute le template avec les données actuelles.
	err = custTemplate.Execute(w, Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DisplayIndex gère la logique d'affichage de la page principale du jeu de pendu.
func DisplayIndex(w http.ResponseWriter, r *http.Request) {
	// Parse le fichier HTML du template.
	custTemplate, err := template.ParseFiles("Hangman-web/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Récupère la lettre entrée par l'utilisateur depuis le formulaire.
	lettre := r.FormValue("lettre")

	// Traite la lettre entrée.
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
						// Redirige vers la page de victoire.
						http.Redirect(w, r, "/win", http.StatusFound)
					} else if Data.Result == "loose" {
						// Redirige vers la page de défaite.
						http.Redirect(w, r, "/loose", http.StatusFound)
					}
				}
			}
		}
		// Met à jour l'image du pendu.
		Data.Jose = Func.DisplayImage(w, r, Data.Tries)
	}

	// Exécute le template avec les données actuelles.
	err = custTemplate.Execute(w, Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DisplayWin gère la logique d'affichage de la page de victoire.
func DisplayWin(w http.ResponseWriter, r *http.Request) {
	// Parse le fichier HTML du template.
	custTemplate, err := template.ParseFiles("Hangman-web/templates/win.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Exécute le template avec les données actuelles.
	err = custTemplate.Execute(w, Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DisplayLoose gère la logique d'affichage de la page de défaite.
func DisplayLoose(w http.ResponseWriter, r *http.Request) {
	// Parse le fichier HTML du template.
	custTemplate, err := template.ParseFiles("Hangman-web/templates/loose.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Exécute le template avec les données actuelles.
	err = custTemplate.Execute(w, Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// StartServer initialise le serveur HTTP pour le jeu de pendu.
func StartServer() {
	println("server started on http://localhost:8080/")
	
	// Gère les chemins pour les fichiers CSS et les images.
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./Hangman-web/templates/css/"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./Hangman-web/assets/images/"))))
	
	// Définit les gestionnaires de routes pour les différentes pages.
	http.HandleFunc("/", Display)
	http.HandleFunc("/index", DisplayIndex)
	http.HandleFunc("/win", DisplayWin)
	http.HandleFunc("/loose", DisplayLoose)
	
	// Lance le serveur sur le port 8080.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
