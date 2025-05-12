/*
1. Crie um tipo. O tipo subjacente deve ser int.
2. Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
3. Na função main:
	1. Demonstre o valor da variável "x"
	2. Demonstre o tipo da variável "x"
	3. Atribua 42 à variável "x" utilizando o operador "="
	4. Demonstre o valor da variável "x"
*/

package main

import "fmt"

type testing int

var x testing

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}