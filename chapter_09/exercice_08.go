/*
Crie um map com key tipo string e value tipo []string.
	Key deve conter nomes no formato sobrenome_nome
	Value deve conter os hobbies favoritos da pessoa
Demonstre todos esses valores e seus índices.
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

	for k, v := range registers {
		fmt.Println(k, v)
	}
}