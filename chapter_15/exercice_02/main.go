/*
Crie um struct "pessoa"
Crie uma função chamada mudeMe que tenha *pessoa como parâmetro.
Essa função deve mudar um valor armazenado no endereço *pessoa.
	Dica: a maneira "correta" para fazer dereference de um valor em um struct seria (*valor).campo
	Mas consta uma exceção na documentação. Link: https://golang.org/ref/spec#Selectors
	"As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method), 
	→ x.f is shorthand for (*x).f." ←
	Ou seja, podemos usar tanto o atalho p1.nome quanto o tradicional (*p1).nome
Crie um valor do tipo pessoa;
Use a função mudeMe e demonstre o resultado.
*/

package main

import "fmt"

type person struct {
	name   string
	access int
}

func main() {
	p1 := person{
		name:   "Daniel",
		access: 0,
	}

	changeMe(&p1)
	changeMe(&p1)
	changeMe(&p1)
	changeMe(&p1)

	fmt.Println("Total de acessos: ", p1.access)
}

func changeMe(p *person) {
	(*p).access += 1
}