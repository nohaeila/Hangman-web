// Package Func fournit une collection de fonctions utilitaires.
package Func

import "net/http"

// DisplayImage prend en paramètres un objet ResponseWriter 'w', un objet Request 'r', et un identifiant d'image 'imageID'.
// La fonction retourne le chemin d'accès à l'image correspondant à l'identifiant fourni.
func DisplayImage(w http.ResponseWriter, r *http.Request, imageID int) string {
	// Initialise la variable 'imagePath'.
	var imagePath string
	
	// Utilise une instruction switch pour déterminer le chemin d'accès à l'image en fonction de l'identifiant 'imageID'.
	switch imageID {
	case 0:
		imagePath = "images/image2.png"
	case 1:
		imagePath = "images/image3.png"
	case 2:
		imagePath = "images/image4.png"
	case 3:
		imagePath = "images/image5.png"
	case 4:
		imagePath = "images/image6.png"
	case 5:
		imagePath = "images/image7.png"
	case 6:
		imagePath = "images/image8.png"
	case 7:
		imagePath = "images/image9.png"
	case 8:
		imagePath = "images/image10.png"
	case 9:
		imagePath = "images/image11.png"
	case 10:
		imagePath = "images/image12.png"
	default:
		imagePath = "images/image1.png"
		
		// Retourne une chaîne vide si l'identifiant n'est pas dans la plage spécifiée.
		return ""
	}
	
	// Retourne le chemin d'accès à l'image correspondant à l'identifiant fourni.
	return imagePath
}
