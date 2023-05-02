package main

import "fmt"

type Triangulo struct {
	Base   float64
	Altura float64
}

func area(t Triangulo) {
	fmt.Println((t.Base * t.Altura) / 2)
}

func main() {
	t := Triangulo{Base: 10, Altura: 20}
	area(t)
}
