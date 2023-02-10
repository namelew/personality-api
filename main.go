package main

import (
	"fmt"

	"github.com/namelew/personality-api/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
