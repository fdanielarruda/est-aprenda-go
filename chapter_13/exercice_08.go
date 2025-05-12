/*
Crie uma função que retorna uma função.
Atribua a função retornada a uma variável.
Chame a função retornada.
*/

package main

import "fmt"

func main() {
	y := master()

	y()
}

func master() func() {
	return func() {
		fmt.Println("Função secundária")
	}
}