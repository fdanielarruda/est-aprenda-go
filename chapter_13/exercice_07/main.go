/*
Atribua uma função a uma variável.
Chame essa função.
*/

package main

import "fmt"

func main() {
	y := func() {
		fmt.Println("Estou chamando essa função em uma variável")
	}

	y()
}