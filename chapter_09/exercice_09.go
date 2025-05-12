/*
Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.
*/

package main

import "fmt"

func main() {
	registers := map[string][]string{
		"arruda_daniel": {
			"Futebol",
			"Tabelas",
			"Xadrez",
		},
		"arruda_alexandre": {
			"Cruzadinhas",
		},
		"arruda_alexsandra": {
			"Flauta",
		},
	}

	registers["vasconcelos_valdelis"] = []string{
		"Terço",
		"Cantar",
	}

	registers["arruda_francisco"] = []string{
		"Baralho",
		"Amigos",
	}

	for key, register := range registers {
		fmt.Println(key)

		for _, hobby := range register {
			fmt.Printf("\t%v\n", hobby)
		}

		fmt.Println("")
	}
}
