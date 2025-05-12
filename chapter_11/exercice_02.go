/*
Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
Demonstre os valores do map utilizando range.
Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
*/

package main

import "fmt"

type person struct {
	name               string
	lastname           string
	favorite_icecreams []string
}

func main() {
	persons := map[string]person{
		"arruda_daniel": person{
			name:     "Daniel",
			lastname: "Arruda",
			favorite_icecreams: []string{
				"Morango",
				"Maracujá",
			},
		},
		"garcia_carol": person{
			name:     "Carol",
			lastname: "Garcia",
			favorite_icecreams: []string{
				"Chocolate",
			},
		},
	}

	for _, person := range persons {
		fmt.Println(person.name, person.lastname)

		for _, icecream := range person.favorite_icecreams {
			fmt.Printf("\t%v\n", icecream)
		}

		fmt.Println("")
	}
}