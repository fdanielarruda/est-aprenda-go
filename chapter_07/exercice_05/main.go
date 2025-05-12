/*
Demonstre o resto da divisão por 4 de todos os números entre 10 e 100
*/

package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("O resto de %v dividido por 4 é: %v\n", i, i % 4)
	}
}