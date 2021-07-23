package main

import (
	"fmt"
)

func main() {
	var TimeAmarelo = []string{"Fernando", "João", "Lúcia", "Mariana", "Ana"}
	var TimeVermelho = []string{"Helena", "Jonas", "José", "Juliana"}
	
	TimeVermelho = append(TimeVermelho, "Luis Inácio")

	fmt.Println("Time Amarelo: ",TimeAmarelo)
	fmt.Println("Time Vermelho: ",TimeVermelho)
}
