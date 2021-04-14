package main

const (
	maxline    int = 3
	maxColonne int = 3
)

var (
	partie  [maxline][maxColonne]int //creation d'un tableau à double dimension d'entier.
	numCase [maxline][maxColonne]int //creation du tableau qui contient les numéros des cases.
)

func afficherPartie() {
	println("Partie en cours : ")
	var compteurCase int = 1 // Sert à afficher le numéro de case
	for i := 0; i < maxline; i++ {
		//Affichage ligne par ligne.
		for j := 0; j < maxColonne; j++ {
			switch partie[i][j] {
			case 3:
				print(compteurCase, " ")
			case 2:
				print("X ")
			case 1:
				print("O ")
			}
			compteurCase++
		}
		println()
	}
}

func Initialize() {
	var numCase int = 1
	for i := 0; i < maxline; i++ {
		for j := 0; j < maxColonne; j++ {
			partie[i][j] = 3
			numCase[i][j] = numCase
			numCase++
		}
	}
}

func gagner() bool {
	var (
		gagne          bool = false //On changera la valeur dès qu'on a une condition trouvée.
		ligne, colonne int  = 0, 1
	)
	// première étape : vérifier qu'il n'y a pas une ligne gagnante.
	for ligne < maxline {
		for partie[ligne][colonne-1] == partie[ligne][colonne] && partie[ligne][colonne] != 3 {
			colonne++
		}
		if colonne == maxColonne {
			gagne = true // On a fait la ligne et on a trouvé la même valeur dans chaque case différent de 3, le joueur a gagné, on retourne true.
		}
		ligne++ // sinon on fait la même chose pour la ligne précédente.
	}

	if !gagne {
		// On arrive ici, on a vérifier qu'aune ligne n'était remplie avec la même valeur. On va faire la même chose mais avec les colonnes.
		ligne, colonne = 1, 0 //Réinitialisation des variables.
		for colonne < maxColonne {
			for partie[ligne-1][colonne] == partie[ligne][colonne] && partie[ligne][colonne] != 3 {
				ligne++
			}
			if ligne == maxline {
				gagne = true // On a fait la colonne et on a trouvé que la colonne était remplie avec la même valeur différente de 3, le joueur a gagné, on retourne true
			}
			colonne++ // sinon, on avance à la colonne suivante.

		}

		if !gagne {
			/*Arrivé ici, aucune ligne et aucune colonne n'est gagnante. On va vérifier la dernière façon de gagner : les deux diagonales.
			Il faut vérifier d'abord que nous sommes dans un carré sinon cela ne fonctionne pas.
			*/
			if maxColonne != maxline {
				gagne = false //différent : on retourne false
			}

			//premier cas : cas où ligne=colonne
			ligne = 1 //réinitialisation.
			for partie[ligne-1][ligne-1] == partie[ligne][ligne] && partie[ligne][ligne] != 3 {
				ligne++
			}

			if ligne == maxline {
				gagne = true // la diagonale est gagnante
			}

			//Deuxième cas : la diagonale opposé celle où on vérifie partie[i+1,j-1] doit être égale à partie[i, j] en commencant par i=0 et j=maxcolonne-1
			ligne = 0
			colonne = maxColonne - 1
			for partie[ligne+1][colonne-1] == partie[ligne][colonne] && partie[ligne][colonne] != 3 {
				ligne++
				colonne--
			}
			if ligne == maxline && colonne == 0 {
				gagne = true // On a parcouru la diagonale et on a eu la même chose sur la diagonale
			}
		}
	}

	return gagne // on retourne gagne qui contient la solution.

}

func menu() {

}

func choixCase(numCase int, joueur int, partie [maxline][maxColonne]int, tabCase [maxline][maxColonne]int) bool {
	//Récupération de l'indice du numéro de case en fonction du numCase
	var (
		ligne, colonne int = 0, 0
		reponse        bool
	)

	for ligne < maxline {
		for colonne < maxColonne {
			if tabCase[ligne][colonne] == numCase {
				break // On arrête la boucle, on a trouvé la ligne et la colonne que l'on voulait.
			}
			colonne++
		}
		ligne++
	}

	if partie[ligne][colonne] == 3 {
		partie[ligne][colonne] = joueur
		reponse = true
	} else {
		println("Case déjà prise, faut faire un autre choix ! ")
		reponse = false
	}
	return reponse
}

func Lapartie() {
	var nbCoupMax int = maxColonne * maxline
	var nbCoup int = 1
	var joueur1, joueur2, joueurActuel string
	var indiceJoueur int
	var reponse int
	//Récupérer le nom des joueurs

	//Init de la partie
	Initialize()

	//Lancement de la partie

	for gagner() || nbCoup <= nbCoupMax {
		if nbCoup%2 == 0 {
			joueurActuel = joueur2
			indiceJoueur = 2
		} else {
			joueurActuel = joueur1
			indiceJoueur = 1
		}
		afficherPartie()
		println(joueurActuel, " Quelle est la case choisie ? ")
		switch {
		case reponse > nbCoupMax || reponse < 0:
			println("choix impossible ! ")
		default:
			if choixCase(reponse, indiceJoueur, partie, numCase) {
				nbCoup++
			}
		}
	}

}

func main() {

}
