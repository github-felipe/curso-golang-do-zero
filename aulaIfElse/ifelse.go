package main

import "fmt"

func main() {
	valor := 1

	if valor == 1 {
		fmt.Println("O valor é realmente 1")
	} else if valor == 2 {
		fmt.Println("O valor é igual a 2")
	} else {
		fmt.Println("O valor é diferente de 1 e 2")
	}

	fmt.Print("Insira um valor: ")
	fmt.Scan(&valor)

	if valor%2 == 0 {
		fmt.Printf("O valor %d é par.\n", valor)
	} else {
		fmt.Printf("O valor %d é ímpar.\n", valor)
	}
}
