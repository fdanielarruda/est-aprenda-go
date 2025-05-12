/*
Crie um programa que utilize a declaração switch, sem switch statement especificado.
*/

package main

import "fmt"

func main() {
	x := 130

	switch {
	case x < 100:
		fmt.Println("X é menor do que 100")
	case x > 200:
		fmt.Println("X é maior do que 200")
	default:
		fmt.Println("X está entre 100 e 200")
	}
}
