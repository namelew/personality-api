package main

import (
	"fmt"

	"github.com/namelew/personality-api/database"
	"github.com/namelew/personality-api/routes"
)

func main() {
	database.Conectar()

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
