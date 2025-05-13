/*
Faça esse código funcionar: https://play.golang.org/p/j-EA6003P0
Usando uma função anônima auto-executável
*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
