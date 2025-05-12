/*
Crie um tipo para um struct chamado "pessoa"
Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
Demonstre no seu código:
	Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
	Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
*/

package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) toSpeak() {
	fmt.Printf("Olá, me chamo %v e tenho %v anos\n", p.name, p.age)
}

type humans interface {
	toSpeak()
}

func saySomething(h humans) {
	h.toSpeak()
}

func main() {
	p1 := person{
		name: "Daniel",
		age:  23,
	}

	p1.toSpeak() // -> é um shortcut para (&p1).toSpeak(). por isso funciona sem passar como ponteiro
	(&p1).toSpeak()

	saySomething(&p1)
	// saySomething(p1) -> não funciona
}
