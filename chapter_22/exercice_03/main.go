/*
Utilizando este código: https://play.golang.org/p/sfyu4Is3FG
...use um for range loop para demonstrar os valores do canal.
*/

package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := range 100 {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
}
