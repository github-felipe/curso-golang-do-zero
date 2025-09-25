package main

import "fmt"

func main() {
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
	fmt.Println(array[0], array[1])
	fmt.Println(array)

	numPrimos := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(numPrimos)
	fmt.Println(numPrimos[2:5])
	fmt.Println(numPrimos[:5])
	fmt.Println(numPrimos[2:])
	fmt.Println(numPrimos[4])
	fmt.Println(numPrimos[2:3])

	slice := make([]string, 2)
	slice[0] = "Hello"
	slice[1] = "World"
	slice = append(slice, "!")
	fmt.Println(slice)

	numPares := []int{2, 4, 6, 8, 10, 12, 14}
	fmt.Println(numPares)
	numPares = append(numPares, 16, 18, 20)
	fmt.Println(numPares)
}
