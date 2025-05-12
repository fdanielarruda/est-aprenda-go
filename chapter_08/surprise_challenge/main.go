/*
Tente acessar todos os itens de uma slice sem utilizar range.
*/

package main

import "fmt"

func main() {
	days := []string{"Segunda", "Terça", "Quarta", "Quinta", "Sexta"}

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
}
