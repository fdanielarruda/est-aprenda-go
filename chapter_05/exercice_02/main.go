/*
Escreva expressões utilizando os operadores, e atribua seus valores a variáveis.
	- ==
	- !=
	- <=
	- <
	- >=
	- >
Demonstre estes valores.
*/

package main

import "fmt"

func main() {
	x := 4
	y := 2

	a := (x == y)
	b := (x != y)
	c := (x <= y)
	d := (x < y)
	e := (x >= y)
	f := (x > y)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}