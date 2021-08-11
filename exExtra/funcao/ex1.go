package main

import (
	"fmt"
)

func main() {
	var num = []int{5,6,89,76,54,3,2,12,34,56,7,8,9,90}
	var soma int = 0

	for x, n := range num {
		if(n % 2 != 0){
			num[x] *= 2
			}

		if(n % 2 == 0){
			num[x] /= 2
		}
		
		soma += num[x]
	}
	fmt.Printf("%v\n%d", num, soma)
}
