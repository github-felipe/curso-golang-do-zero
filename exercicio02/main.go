package main

import "fmt"

func Summation(num int) int {
	value := 0
	for i := 1; i <= num; i++ {
		value += i
	}
	return value
}

func main() {
	fmt.Println(Summation(8))
}
