package main

import "fmt"

func main() {
	var password string
	var reelpass string = "salut"
	var flag string = "G0l4ng_is_34sy"
	fmt.Printf("				Crack Me Golang")
	fmt.Printf("\n\nVeuillez rentrer un mot de passe :  ")
	fmt.Scanf("%s", &password)

	if password == reelpass {
		fmt.Printf("Bien joué, voici le flag {%s}", flag)
	} else {
		fmt.Printf("Réessaye ")
	}
}
