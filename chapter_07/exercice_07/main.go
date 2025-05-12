/*
Utilizando a solução anterior, adicione as opções else if e else.
*/

package main

import "fmt"

func main() {
	x := 10
	y := 15

	if x > y {
		fmt.Println("x é maior que y")
	} else if x == y {
		fmt.Println("x é igual a y")
	} else {
		fmt.Println("x é menor que y")
	}
}