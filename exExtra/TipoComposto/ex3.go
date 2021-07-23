package main

import (
	"fmt"
	
)

func main() {
cidade := map[string]string{
"Arcoverde": "Brasil",
"São Paulo": "Brasil",
"Recife": "Brasil",
"Boston": "EUA",
"Montreal": "Canada",
"Berlin": "Alemanha",
}

frequencia := make(map[string]int)
for _, valor := range cidade{
		frequencia[valor] += 1
	}
	
fmt.Println(frequencia)
	


}
