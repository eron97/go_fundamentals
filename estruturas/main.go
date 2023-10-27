package main

import "fmt"

func main() {

	// Criação de uma struct
	type Person struct {
		age  int
		name string
	}

	// Criação de uma instância da struct
	var person Person

	// Atribuição de valores
	person.age = 30
	person.name = "John Doce"

	fmt.Printf("Name: %s with age: %d", person.name, person.age)
	// saída -> Name: John Doce with age: 30

}
