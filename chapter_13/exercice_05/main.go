/*
Crie um tipo "quadrado"
Crie um tipo "círculo"
Crie um método "área" para cada tipo que calcule e retorne a área da figura
	Área do círculo: 2 * π * raio
	Área do quadrado: lado * lado
Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
Crie um valor de tipo "quadrado"
Crie um valor de tipo "círculo"
Use a função "info" para demonstrar a área do "quadrado"
Use a função "info" para demonstrar a área do "círculo"
*/

package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return 2 * math.Pi * c.radius
}

type figure interface {
	area() float64
}

func info(f figure) {
	fmt.Println(f.area())
}

func main() {
	s1 := square{
		side: 10,
	}

	c1 := circle{
		radius: 5,
	}

	info(s1)
	info(c1)
}
