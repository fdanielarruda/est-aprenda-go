/*
Crie um struct "erroEspecial" que implemente a interface builtin.error.
Crie uma função que tenha um valor do tipo error como parâmetro.
Crie um valor do tipo "erroEspecial" e passe-o para a função da instrução anterior.
*/

package main

import "fmt"

type CustomError struct {
	details string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("Custom error occurred: %v", e.details)
}

func handleError(err error) {
	if customErr, ok := err.(CustomError); ok {
		fmt.Println("Internal error message:", customErr.details)
		fmt.Println("Error() method output:", err)
	} else {
		fmt.Println("Generic error:", err)
	}
}

func main() {
	customErr := CustomError{"Something went wrong with the system"}
	handleError(customErr)
}
