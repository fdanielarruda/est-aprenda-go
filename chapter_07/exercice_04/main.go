/*
Crie um loop utilizando a sintaxe: for {}
Utilize-o para demonstrar os anos desde que você nasceu.
*/

package main

import "fmt"

func main() {
	start_year := 2002
	current_year := 2025

	for {
		if start_year > current_year {
			break
		}

		fmt.Println(start_year)
		start_year++
	}
}