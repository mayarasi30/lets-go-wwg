package main

import "fmt"

func main(){

	var valor int
	
	fmt.Println("Digite o primeiro numero:")
	fmt.Scanln(&valor)
	
switch  {
	case valor == 0:
		fmt.Println("O número é 0.")
	case valor % 2 == 0:
		fmt.Println("O número é múltiplo de 2.")
	case valor % 3 == 0:
		fmt.Println("O número é múltiplo de 3.")
	default:
		fmt.Println("O número não é 0 nem multiplo de 2 ou de 3.")

	}
}

