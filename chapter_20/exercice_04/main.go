/*
Utilize mutex para consertar a condição de corrida do exercício anterior.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var count int

const amountOfRoutines = 10

func main() {
	newRoutines(amountOfRoutines)
	wg.Wait()

	fmt.Println("Total de go routines: ", amountOfRoutines)
	fmt.Println("Total do contador: ", count)
}

func newRoutines(i int) {
	wg.Add(i)

	for range i {
		go func() {
			mu.Lock()
			c := count
			runtime.Gosched()
			c++
			count = c
			wg.Done()
			mu.Unlock()
		}()
	}
}
