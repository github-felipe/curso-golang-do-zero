package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

type Programador struct {
	Pessoa
	Linguagem string
}

func main() {
	fmt.Println(Pessoa{"Felipe", 22})
	fmt.Println(Pessoa{Nome: "Felipe"})

	p1 := Pessoa{Nome: "Felipe", Idade: 22}
	fmt.Println(p1)
	fmt.Println(p1.Nome)
	fmt.Println(p1.Idade)

	p1.Idade = 23
	fmt.Println(p1.Idade)

	p2 := Pessoa{Nome: "Robson"}
	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)
	fmt.Println(pessoas)

	Moradores := map[string][]Pessoa{}
	Moradores["Araraquara"] = pessoas

	fmt.Println(Moradores)
	fmt.Println(Moradores["Araraquara"])
	fmt.Println(Moradores["Araraquara"][0])
	fmt.Println(Moradores["Araraquara"][0].Nome)

	var alunos = map[string][]Pessoa{
		"Programação": {{Nome: "Felipe", Idade: 22}, {Nome: "Steph", Idade: 28}},
		"Engenharia":  {{Nome: "Felipe2", Idade: 22}, {Nome: "Steph2", Idade: 28}},
	}
	fmt.Println(alunos)
	fmt.Printf("Alunos de Programação: %v\n", alunos["Programação"])
	fmt.Printf("Alunos de Engenharia: %v\n", alunos["Engenharia"])

	p1.Idade = 22
	prog := Programador{Pessoa: p1, Linguagem: "Golang"}

	fmt.Println(prog)
	fmt.Println(prog.Pessoa)
	fmt.Println(prog.Pessoa.Nome)
	fmt.Println(prog.Pessoa.Idade)
	fmt.Println(prog.Linguagem)
}
