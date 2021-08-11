package main

import (
	"fmt"
)



func main() {
txt:="litro"
fmt.Println(txt)
t1:=add3(txt,3)
fmt.Println(t1)
fmt.Println(add3(t1,-3))
}

func add3(texto string, num int) string {
	
	var novoTexto string
	for i := 0; i < len(texto); i++ {
	novoTexto += string(texto[i]+byte(num))
	}
	return(novoTexto)
}



