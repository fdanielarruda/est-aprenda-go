/*
Crie um tipo struct "pessoa" que contenha os campos:
	nome
	sobrenome
	idade
Crie um método para "pessoa" que demonstre o nome completo e a idade;
Crie um valor de tipo "pessoa";
Utilize o método criado para demonstrar esse valor.
*/

package main

import "fmt"

type person struct {
	name     string
	lastname string
	age      int
}

func (p person) mySelf() {
	fmt.Printf("Me chamo %v %v e tenho %v anos.", p.name, p.lastname, p.age)
}

func main() {
	person1 := person{
		name:     "Daniel",
		lastname: "Arruda",
		age:      23,
	}

	person1.mySelf()
}