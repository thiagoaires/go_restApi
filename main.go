package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Pizza struct {
	ID    int
	nome  string
	preco float64
}

func main() {
	router := gin.Default()

	router.GET("/pizzas", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"pizzas": "Toscana, Calabresa",
		})
	})

	router.Run()

	var pizzas = []Pizza{
		{ID: 1, nome: "Toscana", preco: 12.94},
		{ID: 2, nome: "Toscana", preco: 65.12},
		{ID: 3, nome: "Toscana", preco: 123.94},
	}
	fmt.Print(pizzas)
}
