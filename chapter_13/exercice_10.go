/*
Demonstre o funcionamento de um closure.
Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.
*/

package main

import "fmt"

func main() {
	a := primary()
	a()
	a()
	a()

	fmt.Println("")

	b := primary()
	b()
	b()
	b()
}

func primary() func() {
	i := 0

	return func() {
		i++
		fmt.Println(i)
	}
}
