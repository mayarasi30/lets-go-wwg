package main

import "fmt"

func main(){

	var v1,v2,v3 int
	
	fmt.Println("Digite o primeiro numero:")
	fmt.Scanln(&v1)
	
	fmt.Println("Digite o segundo numero:")
	fmt.Scanln(&v2)
	
	fmt.Println("Digite o terceiro numero:")
	fmt.Scanln(&v3)

	if v3 > v2 && v3 > v1{
		fmt.Println("o maior valor e:",v3)
	}else if v2 > v3 && v2 > v1{
		fmt.Println("o maior valor e:",v2)
	}else if v1 > v3 && v1 > v1{
		fmt.Println("o maior valor e:",v1)
	}

}
