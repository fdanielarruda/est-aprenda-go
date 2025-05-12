/*
Callback: passe uma função como argumento a outra função.
*/

package main

import "fmt"

func main() {
	sum(spanish, 10, 15, 24)
}

func sum(f func(int), n ...int) {
	total := 0

	for _, v := range n {
		total += v
	}

	f(total)
}

func portuguese(x int) {
	fmt.Println("O resultado dessa operação é", x)
}

func english(x int) {
	fmt.Println("The result of this operation is", x)
}

func spanish(x int) {
	fmt.Println("El resultado de esta operación es", x)
}