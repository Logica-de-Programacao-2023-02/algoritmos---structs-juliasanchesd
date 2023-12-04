package main

import "fmt"

type Triangulo struct {
	base   int
	altura int
}

func main() {
	p := Triangulo{base: 2, altura: 4}
	a := calc(p)
	fmt.Print(a)
}

func calc(p Triangulo) int {
	return p.base * p.altura / 2
}
