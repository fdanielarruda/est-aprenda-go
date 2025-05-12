/*
Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
Passe um valor do tipo slice of int como argumento para a função;
Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
Passe um valor do tipo slice of int como argumento para a função.
*/

package main

import "fmt"

func main() {
	value := []int{10, 15, 20, 25, 30}

	total := sum(value...)
	fmt.Println(total)

	total2 := sum2(value)
	fmt.Println(total2)
}

func sum(x ...int) int {
	total := 0

	for _, v := range x {
		total += v
	}

	return total
}

func sum2(x []int) int {
	total := 0

	for _, v := range x {
		total += v
	}

	return total
}