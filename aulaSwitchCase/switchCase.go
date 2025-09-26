package main

import "fmt"

func main() {
	var posicao int
	fmt.Print("Insira a posição em que você ficou no ranking: ")
	fmt.Scan(&posicao)

	switch posicao {
	case 1:
		fmt.Println("Parabéns, você ganhou!")

	case 2:
		fmt.Println("Parabéns pelo segundo lugar!")

	case 3:
		fmt.Println("Parabéns pelo terceiro lugar!!!")

	default:
		fmt.Println("Ao menos valeu a participação!")
	}
}
