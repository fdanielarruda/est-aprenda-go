/*
Crie um loop utilizando a sintaxe: for condition {}
Utilize-o para demonstrar os anos desde que você nasceu.
*/

package main

import "fmt"

func main() {
	start_year := 2002
	current_year := 2025

	for start_year <= current_year {
		fmt.Println(start_year)
		start_year++
	}
}