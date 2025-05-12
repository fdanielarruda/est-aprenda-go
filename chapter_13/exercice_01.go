/*
Crie uma função que retorne um int
Crie outra função que retorne um int e uma string
Chame as duas funções
Demonstre seus resultados
*/

package main

import "fmt"

func main() {
	fmt.Println(first())
	fmt.Println(second())
}

func first() int {
	return 1
}

func second() (int, string) {
	return 2, "b"
}