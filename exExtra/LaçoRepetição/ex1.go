package main

import (
	"fmt"
)

func main() {
	Mercado := []string{"Banana", "Arroz", "Feijão", "Tomate", "Frango", "Acucar"}

	for lista := 0; lista < 6; lista++{
		fmt.Printf("%d %s\n", lista, Mercado[lista])
	}
}
