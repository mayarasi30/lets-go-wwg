package main

import (
	"fmt"
	"os"
)

func main() {
	var num
	var x
	for x == 0{
		fmt.Println("Insira um número:")
		fmt.Fscanf(os.Stdin, "%d", &num)
		if num % 2 == 0{
			n=n+1
		}
	}
	fmt.Println("Agora sim podemos dividir igualmente entre mim e você!")
}
