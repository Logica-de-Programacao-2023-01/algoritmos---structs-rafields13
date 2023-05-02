package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	Raio float64
}

func area(c Circulo) {
	fmt.Println(math.Pi * math.Pow(c.Raio, 2))
}

func main() {
	c := Circulo{Raio: 13.7}
	area(c)
}
