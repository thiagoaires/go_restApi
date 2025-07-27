package main

import "fmt"

type Pizza struct {
	ID    int
	nome  string
	preco float64
}

func main() {
	// toscana, marguerita, atum com queijo
	// nomePizzaria := "Pizza a goGO"
	// insta, phone := "@pizza_go", 11911111111
	// fmt.Println(nomePizzaria, insta, phone)

	var pizzas = []Pizza{
		{ID: 1, nome: "Toscana", preco: 12.94},
		{ID: 2, nome: "Toscana", preco: 65.12},
		{ID: 3, nome: "Toscana", preco: 123.94},
	}
	fmt.Print(pizzas)
}
