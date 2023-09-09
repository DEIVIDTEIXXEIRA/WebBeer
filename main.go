package main

import (
	"fmt"
	"webbeer/config"
)

func main() {
	config.Carregar()
	fmt.Printf("Escutando na porta: %d", config.Porta)
}
