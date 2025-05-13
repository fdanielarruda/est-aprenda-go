/*
Crie um programa que lance 10 goroutines onde cada uma envia 10 números a um canal;
Tire estes números do canal e demonstre-os.
*/

package main

import "fmt"

func main() {
	channel := make(chan int)

	gen(channel, 10)
	read(channel)
}

func gen(c chan int, n int) {
	for range n {
		go func() {
			for j := range n {
				c <- j
			}
		}()
	}
}

func read(c chan int) {
	for k := range 100 {
		fmt.Println(k, <-c)
	}
}
