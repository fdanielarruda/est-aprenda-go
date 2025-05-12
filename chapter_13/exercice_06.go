/*
Crie e utilize uma função anônima.
*/

package main

import "fmt"

func main() {
	func() {
		fmt.Println("Olá, meus amigos!")
	}()
}