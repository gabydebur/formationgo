package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Pour le moment, je le mets en constante parce que j'ai pas encore vu les tableaux dynamique ! :p
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
			//println("partie i et j :", i, j, partie[i][j])
			switch partie[i][j] {
			case 3:
				//println("pas de joueur'")
				print(compteurCase, " ")
			case 2:
				//println("joueur 2")
				print("X ")
			case 1:
				//println("joueur1")
				print("O ")
			}
			compteurCase++
		}
		println()
	}
}

func Initialize() {
	var numeroCase int = 1
	var i, j int
	for i = 0; i < maxline; i++ {
		for j = 0; j < maxColonne; j++ {
			partie[i][j] = 3
			numCase[i][j] = numeroCase
			numeroCase++
			print("-")
		}
	}
	println("")
}

// un peu complexe, a diviser en 3 fonctions par exemple checkWinLine, checkWinColumn, checkWinDiagonale
func gagner() bool {
	var (
		gagne          bool = false //On changera la valeur dès qu'on a une condition trouvée.
		//perso j'aurais init les deux a 0 et j'aurais fait +1 au lieu de -1 dans les for, c'est plus comprehensible je pense
		ligne, colonne int  = 0, 1
	)
	// première étape : vérifier qu'il n'y a pas une ligne gagnante.
	for ligne < maxline && !gagne {
		for colonne < maxColonne && partie[ligne][colonne-1] == partie[ligne][colonne] && partie[ligne][colonne] != 3 {
			colonne++
			//println("colonne++", colonne)
		}
		if colonne == maxColonne {
			// return true , comme ça tu peux supprimer les if !gagne en dessous, ça simplifie le code
			gagne = true // On a fait la ligne et on a trouvé la même valeur dans chaque case différent de 3, le joueur a gagné, on retourne true.
		} else {
			ligne++ // sinon on fait la même chose pour la ligne suivante.
		}
	}
	// a supprimer cf commentaire au dessus
	if !gagne {
		// On arrive ici, on a vérifier qu'aune ligne n'était remplie avec la même valeur. On va faire la même chose mais avec les colonnes.
		ligne, colonne = 1, 0 //Réinitialisation des variables.
		for colonne < maxColonne && !gagne {
			for ligne < maxline && partie[ligne-1][colonne] == partie[ligne][colonne] && partie[ligne][colonne] != 3 {
				ligne++
			}
			if ligne == maxline {
				// return true , comme ça tu peux supprimer les if !gagne en dessous, ça simplifie le code
				gagne = true // On a fait la colonne et on a trouvé que la colonne était remplie avec la même valeur différente de 3, le joueur a gagné, on retourne true
			} else {
				colonne++ // sinon, on avance à la colonne suivante.
			}
		}
		// a supprimer cf commentaire au dessus
		if !gagne {
			/*Arrivé ici, aucune ligne et aucune colonne n'est gagnante. On va vérifier la dernière façon de gagner : les deux diagonales.
			Il faut vérifier d'abord que nous sommes dans un carré sinon cela ne fonctionne pas.
			*/
			if maxColonne != maxline {
				gagne = false //différent : on retourne false
			} else {
				//premier cas : cas où ligne=colonne
				ligne = 1 //réinitialisation.
				for ligne < maxline && partie[ligne-1][ligne-1] == partie[ligne][ligne] && partie[ligne][ligne] != 3 {
					ligne++
				}

				if ligne == maxline {
					gagne = true // la diagonale est gagnante
				} else {
					//Deuxième cas : la diagonale opposé celle où on vérifie partie[i+1,j-1] doit être égale à partie[i, j] en commencant par i=0 et j=maxcolonne-1
					ligne = 0
					colonne = maxColonne - 1
					for ligne < maxline-1 && partie[ligne+1][colonne-1] == partie[ligne][colonne] && partie[ligne][colonne] != 3 {
						ligne++
						colonne--
					}
					if ligne == maxline-1 && colonne == 0 {
						gagne = true // On a parcouru la diagonale et on a eu la même chose sur la diagonale
					}
				}
			}
		}
	}
	// ici ça devient juste return false, vu qu'a chaque fois que la condition de victoire etait remplis on aura return true directement
	//	println("On retourne si la partie est gagné ou pas ", gagne)
	return gagne // on retourne gagne qui contient la solution.

}

func choixCase(numeroCase int, joueur int) bool {
	//Récupération de l'indice du numéro de case en fonction du numCase
	var (
		ligne, colonne int = 0, 0
		reponse        bool
		trouve         bool = false
	)
	//println("numcase", numeroCase, " et indice joueur :", joueur)
	for ligne < maxline && !trouve {
		//println("ligne : ", ligne)
		colonne = 0
		for colonne < maxColonne && !trouve {
			//println("colone: ", colonne)
			if numCase[ligne][colonne] == numeroCase {
				//println("jai trouvé !")
				trouve = true
			} else {
				colonne++
			}
		}
		if !trouve {
			ligne++
		}
	}

	//println("Après la boucle, ligne et colonne : ", ligne, colonne)
	if trouve {
		if partie[ligne][colonne] == 3 {
			//println("Case pas prise ! affectation de la case pour le joueur")
			partie[ligne][colonne] = joueur
			//println("ligne, colonne, case : ", ligne, " ", colonne, " ", partie[ligne][colonne])
			reponse = true
		} else {
			println("Case déjà prise, faut faire un autre choix ! ")
			reponse = false
		}
	} else {
		println("Pas trouvé la case ! ")
		reponse = false
	}
	return reponse
}

func Lapartie() {
	var nbCoupMax int = maxColonne * maxline
	var nbCoup int = 1
	var joueur1, joueur2, joueurActuel string
	var indiceJoueur int

	//Récupérer le nom des joueurs
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Joueur 1, Entrez votre nom : ")
	scanner.Scan()
	joueur1 = scanner.Text()
	fmt.Print("Joueur 2, Entrez votre nom : ")
	scanner.Scan()
	joueur2 = scanner.Text()

	//Init de la partie
	println("Initialisation de la partie ! ")
	Initialize()
	println("Initialisation complète!! Début de la partie !!! Bon amusement !")

	//Lancement de la partie

	for !gagner() && nbCoup <= nbCoupMax {
		if nbCoup%2 == 0 {
			joueurActuel = joueur2
			indiceJoueur = 2
		} else {
			joueurActuel = joueur1
			indiceJoueur = 1
		}
		afficherPartie()
		println(joueurActuel, " Quelle est la case choisie ? ")
		scanner.Scan()
		reponse, err := strconv.Atoi(scanner.Text())
		if err != nil {
			println("Ce n'est pas un choix valide !!!")
		} else {
			switch {
			case reponse > nbCoupMax || reponse < 0:
				println("choix impossible ! ")
			default:
				if choixCase(reponse, indiceJoueur) {
					nbCoup++
				}
			}
		}
	}
	afficherPartie()
	if gagner() {
		println(joueurActuel, " a gagné !! Bravo à lui ! !!")
	} else {
		println("Partie nulle ! Ce fut un beau match ! ")
	}
}

func afficherMenuEtChoix() int {
	scanner := bufio.NewScanner(os.Stdin)
	var resultat int
	println("MENU : ")
	println("1 => Lancement de la partie")
	println("2 => Quitter ! ")
	println("Quelle est votre choix ? ")
	scanner.Scan()
	choix, err := strconv.Atoi(scanner.Text())
	if err != nil {
		resultat = 0
	} else {
		resultat = choix
	}
	return resultat

}

func main() {
	println("Bienvenue au jeu du morpion by Gaby")
	choix := afficherMenuEtChoix()
	for choix != 2 {
		switch choix {
		case 1:
			Lapartie()
			choix = afficherMenuEtChoix()
		default:
			println("Choix invalide!! Faire un autre choix ! ")
			choix = afficherMenuEtChoix()
		}
	}
	println("Au Revoir !!! :) ")
}
