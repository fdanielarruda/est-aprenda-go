/*
Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
	Nome
	Sobrenome
	Sabores favoritos de sorvete
Crie dois valores do tipo "pessoa" e demonstre estes valores,
utilizando range na slice que contem os sabores de sorvete.
*/

package main

import "fmt"

type person struct {
	name               string
	lastname           string
	favorite_icecreams []string
}

func main() {
	person1 := person{
		name:     "Daniel",
		lastname: "Arruda",
		favorite_icecreams: []string{
			"Morango",
			"Maracujá",
		},
	}

	fmt.Println("Nome:", person1.name, person1.lastname)
	fmt.Print("Sorvetes favoritos: ")
	for _, icecream := range person1.favorite_icecreams {
		fmt.Print(icecream, " ")
	}
}
