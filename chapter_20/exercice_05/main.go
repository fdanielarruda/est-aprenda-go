/*
Utilize atomic para consertar a condição de corrida do exercício #3.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var count int64

const amountOfRoutines = 100

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
			atomic.AddInt64(&count, 1)
			v := atomic.LoadInt64(&count)
			runtime.Gosched()
			fmt.Println(v)
			wg.Done()
		}()
	}
}
