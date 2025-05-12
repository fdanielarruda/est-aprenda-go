/*
Crie uma slice contendo slices de strings ([][]string).
Atribua valores a este slice multi-dimensional da seguinte maneira:
	"Nome", "Sobrenome", "Hobby favorito"
Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
*/

package main

import "fmt"

func main() {
	registers := [][]string{
		[]string{
			"Daniel",
			"Arruda",
			"Futebol",
		},
		[]string{
			"Alexandre",
			"Arruda",
			"Cruzadinhas",
		},
		[]string{
			"Alexsandra",
			"Arruda",
			"Flauta",
		},
	}

	for _, register := range registers {
		fmt.Printf("Nome: %v\tSobrenome: %v\tHobby: %v\n", register[0], register[1], register[2])
	}
}
