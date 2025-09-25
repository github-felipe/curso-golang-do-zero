package main

import "fmt"

func main() {
	fmt.Println(soma(42, 13))
	fmt.Println(subtracao(42, 13))

	nome1, nome2, _ := retornaNome("Felipe")
	fmt.Println(nome1)
	fmt.Println(nome2)

	nome3, _, _ := retornaNome("felipe")
	fmt.Println(nome3)

	nomeCompleto := retornaNomeCompleto("Felipe", "Lima")
	fmt.Println(nomeCompleto)
}

func soma(x int, y int) int {
	return x + y
}

func subtracao(x int, y int) int {
	return x - y
}

func retornaNome(nome string) (string, string, string) {
	return nome, nome, nome
}

func retornaNomeCompleto(nome, sobrenome string) string {
	return nome + " " + sobrenome
}
