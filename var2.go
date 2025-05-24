package main

import "fmt"

type Pessoa struct {
	Nome     string
	Idade    uint8
	RG       int
	Endereco string
	CEP      int
}

func main() {

	p := Pessoa{
		Nome:     "Marlon Costa",
		Idade:    33,
		RG:       881112223,
		Endereco: "Rua Gavea, 71 - PG",
		CEP:      01111111,
	}

	fmt.Println(p)

}
