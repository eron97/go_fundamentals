package main

import (
	"fmt"
	"math"
)

// Definindo a interface Shape
type Shape interface {
	Area() float64
}

// Definindo a estrutura Circle
type Circle struct {
	Radius float64
}

// Implementação do método Area para Circle
// Com Função acompanhado de receptor (receiver function):
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Definindo a estrutura Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Implementação do método Area para Rectangle
// Com Função sem receptor (receiver function):
func Area(r Rectangle) float64 {
	return r.Width * r.Height
}

func main() {

	// Criando uma instância de Circle
	circle := Circle{Radius: 5}

	// Criando uma instância de Rectangle
	rectangle := Rectangle{Width: 4, Height: 6}

	// Calculando a área do círculo usando a função name_function

	fmt.Printf("Área do círculo: %.2f\n", circle.Area())
	// saída -> Área do círculo: 78.54

	fmt.Printf("Área do retângulo: %.2f\n", Area(rectangle))
	// saída -> Área do retângulo: 24.00

}
