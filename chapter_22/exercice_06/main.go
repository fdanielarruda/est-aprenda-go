/*
Escreva um programa que coloque 100 números em um canal, retire os números do canal, e demonstre-os.
*/

package main

import "fmt"

func main() {
	channel := make(chan int)

	gen(channel)
	read(channel)
}

func gen(c chan int) {
	go func() {
		for i := range 100 {
			c <- i + 1
		}

		close(c)
	}()
}

func read(c chan int) {
	for range 100 {
		fmt.Println(<-c)
	}
}
