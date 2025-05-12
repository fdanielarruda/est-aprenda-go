/*
Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
Demonstre estes valores.
*/

package main

import "fmt"

const (
	_ = 2024 + iota
	a
	b
	c
	d
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
