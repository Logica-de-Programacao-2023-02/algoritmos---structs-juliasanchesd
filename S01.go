package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func main() {
	p := Circulo{raio: 2}
	a := area(p)
	fmt.Print(a)
}

func area(p Circulo) float64 {
	return math.Pi * p.raio * p.raio
}
