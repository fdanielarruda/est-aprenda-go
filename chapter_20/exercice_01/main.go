/*
Alem da goroutine principal, crie duas outras goroutines.
Cada goroutine adicional devem fazer um print separado.
Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	newRoutines(2)
	wg.Wait()
}

func newRoutines(i int) {
	wg.Add(i)

	for j := 1; j <= i; j++ {
		go func(i int) {
			fmt.Println("GO ROUTINE Nº", i)
			wg.Done()
		}(j)
	}
}
