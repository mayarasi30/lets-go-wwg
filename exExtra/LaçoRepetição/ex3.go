package main

import (
	"fmt"
)


func main() {
for x := 1; x < 10; x++ {
		  for i := 1; i < x; i++ {
		    fmt.Printf("%d",i)
		  }
	fmt.Printf("%d \n",x)	
	}

}
