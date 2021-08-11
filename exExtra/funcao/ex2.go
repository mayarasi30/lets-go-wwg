package main

import (
	"fmt"
	"strings"
)

func main() {
	texto := "Programar é bom mas não é fácil não"
	letra := "r"
	s := ContLetra(texto,letra)
	fmt.Printf("A letra %s foi encontrada %d vezes na frase '%s'", letra, s, texto)
	
}
	
func ContLetra(t string, l string) int {


	var final = strings.Count(t, l)
	
	return final
}
