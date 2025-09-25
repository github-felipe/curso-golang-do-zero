package main

import "fmt"

func main() {
	var nome string

	nome = "Felipe"
	fmt.Println("Nome: " + nome)
	
	nome = "Arnaldo"
	fmt.Println("Nome:", nome)

	var a = "steph"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	d := "linux"
	fmt.Println(d)

	const num = 5
	fmt.Println(num)
}