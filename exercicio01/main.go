package main

import "fmt"

func MakeNegative(numero int) int {
	if numero < 0 {
		return numero
	} else {
		return -numero
	}
}

func main() {

	fmt.Println(MakeNegative(1))
	fmt.Println(MakeNegative(-1))
	fmt.Println(MakeNegative(999))
	fmt.Println(MakeNegative(-2345))
	fmt.Println(MakeNegative(0))
}
