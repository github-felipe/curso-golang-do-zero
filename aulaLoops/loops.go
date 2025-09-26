package main

import (
	"fmt"
)

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		fmt.Println("Soma:", sum)
		fmt.Println("Índice:", i)
		sum += i
	}
	fmt.Println(sum)

	frutas := []string{"Maçã", "Banana", "Uva", "Laranja", "Abacaxi"}

	for _, fruta := range frutas {
		fmt.Println(fruta)
	}
}
