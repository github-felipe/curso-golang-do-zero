package main

import "fmt"

func main() {
	idade := map[string]int{}
	idade["Felipe"] = 22
	idade["Robson"] = 34
	idade["Steph"] = 28
	idade["Bento"] = 4

	fmt.Println(idade)
	fmt.Println("Idade do Felpe:", idade["Felipe"])

	anoNasc := map[string]int{
		"Steph":  1995,
		"Bento":  2008,
		"Felipe": 2003,
	}
	anoNasc["GolangDoZero"] = 2024
	fmt.Println(anoNasc)
}
