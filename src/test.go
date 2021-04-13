package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	/*
		pour commenter sur
		plusieurs ligne
		s de texte
	*/
	var (
		vie               int     = 32
		argent, puissance int     = 20, 30
		nom               string  = "gaby"
		vitesse           float32 = 6.4
	)
	fmt.Println("Hello, World!") // pour afficher du texte !
	fmt.Println("vie du coup est égale à", vie)
	fmt.Println("puissance et argent : ", argent, "et ", puissance)
	fmt.Println("nom", nom, "et vitesse", vitesse)

	/*
	   Déclaration des variables dynamiques
	*/
	flt := 15.5   //  sera automatiquement de type float
	in := 5       //  sera automatiquement de type int
	st := "hello" //  sera automatiquement de type string
	bol := true   //  sera automatiquement de type boolean

	fmt.Printf("Le type de la varialbe flt est %T\n", flt)
	fmt.Printf("Le type de la varialbe in est %T\n", in)
	fmt.Printf("Le type de la varialbe st est %T\n", st)
	fmt.Printf("Le type de la varialbe bol est %T\n", bol)

	const maConstante int = 50 // déclaration d'une constante

	fmt.Println("ma Constante : ", maConstante)

	var a int = 4
	var b int = 2

	fmt.Println("a + b  = ", a+b) // addition de la variable a et b
	fmt.Println("a - b  = ", a-b) // soustraction de la variable a et b
	fmt.Println("a * b  = ", a*b) // multiplication de la variable a et b
	fmt.Println("a / b  = ", a/b) // division de la variable a et b
	fmt.Println("a % b  = ", a%b) // modulo de la variable a et b

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Votre choix : ")
	scanner.Scan()
	choix, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Entrez un entier !")
		os.Exit(2)
	}

	switch {
	case choix >= 2000:
		println("Ah un 2000 !")
	case choix >= 1939 && choix <= 1945:
		println("Triste année")
	case time.Now().Weekday() == time.Sunday:
		println("Dimanche !")
	default:
		println("Mauvais choix !") // Valeur par défaut
	}

	fmt.Print("Entrez votre age : ")
	scanner.Scan()
	age, err := strconv.Atoi(scanner.Text())
	if err != nil {
		// Si la conversion n'a pas fonctionné alors on affiche un message d'erreur et on quitte le programme
		fmt.Println("On essaie de m'arnaquer ? allez Dehors ! Et la prochaine entrez un entier !")
		os.Exit(2) // on quitte le programmation
	}

	if age < 17 { // vérifier si l'utilisateur à au moins 18 ans
		fmt.Println("Sortez !")
	} else { // si ce n'est pas le cas alors on l'accepte pas
		fmt.Println("Entrez :)")
	}
}
