/*
Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.
*/

package main

import "fmt"

func main() {
	registers := map[string][]string{
		"arruda_daniel": {
			"assitir futebol",
			"fazer tabelas",
			"jogar xadrez",
		},
		"messi_lionel": {
			"brincar com os filhos",
			"assistir o barcelona",
		},
	}

	delete(registers, "messi_lionel")

	for key, hobbies := range registers {
		fmt.Println(key)

		for _, hobby := range hobbies {
			fmt.Printf("\t%v\n", hobby)
		}

		fmt.Println("")
	}
}
