/*
Crie e use um struct anônimo.
Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
*/

package main

import "fmt"

func main() {
	model := struct {
		independencia map[string]int
		gosma         []string
	}{
		independencia: map[string]int{
			"brasil": 1822,
			"chile":  1818,
		},
		gosma: []string{
			"vasco", "gama",
		},
	}

	fmt.Println(model)
}
