package main

import (
	"fmt"
	"time"
)

func main(){

	var nascimento int
	fmt.Printf("Em que ano você nasceu: ")
	fmt.Scanln(&nascimento)
	
	data := time.Now()
	ano:= data.Year()
	
	idade := ano - nascimento

	fmt.Println("Sua Idade e: ", idade)
	
}
