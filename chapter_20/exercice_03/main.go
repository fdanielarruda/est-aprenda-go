/*
Utilizando goroutines, crie um programa incrementador:
	Tenha uma variável com o valor da contagem
	Crie um monte de goroutines, onde cada uma deve:
		Ler o valor do contador
		Salvar este valor
		Fazer yield da thread com runtime.Gosched()
		Incrementar o valor salvo
		Copiar o novo valor para a variável inicial
	Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
	Demonstre que há uma condição de corrida utilizando a flag -race
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
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
			c := count
			runtime.Gosched()
			c++
			count = c

			wg.Done()
		}()
	}
}
