/*
Faça esse código funcionar: https://play.golang.org/p/j-EA6003P0
Usando buffer
*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
