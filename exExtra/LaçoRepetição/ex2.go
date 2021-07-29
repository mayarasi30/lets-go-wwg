package main

import (
	"fmt"
)
type Apt struct {
	numero        int
	Proprietario  string
	Garagem       bool
}

func main() {
	ap1 := Apt{
		numero:             567,
		Proprietario:  "Mayara",
		Garagem:           true,
	}

	ap2 := Apt{
		numero:             500,
		Proprietario:  "Thayna",
		Garagem:          false,
	}

	ap3 := Apt{
		numero:              57,
		Proprietario:   "Thays",
		Garagem:           true,
	}
	ap4 := Apt{
		numero:              67,
		Proprietario:   "Erick",
		Garagem:          false,
	}

	apartamentos := []Apt{ap1, ap2, ap3, ap4}

	for _, apartamento := range apartamentos {
		possuiVaga := "Não"
		if apartamento.Garagem {
			possuiVaga = "Sim"
		}
		fmt.Printf("O apartamento número %d tem como proprietária %s. Ele tem vaga de garagem? %s\n", apartamento.numero, apartamento.Proprietario, possuiVaga)
	}

}
