/*
Crie um programa que utilize a declaração switch,
onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".
*/

package main

import "fmt"

func main() {
	esporteFavorito := "futebol"

	switch esporteFavorito {
	case "futebol":
		fmt.Println("O melhor time de futebol é o vasco da gama")
	case "basquete":
		fmt.Println("O melhor time de basquete é o chicago bulls")
	case "formula 1":
		fmt.Println("A melhor equipe de fórmula 1 é a red bull")
	default:
		fmt.Println("Eu não gosto de esportes")
	}
}