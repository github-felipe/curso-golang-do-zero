package main

import "fmt"

func CalculateYears(years int) (result [3]int) {
	result[0] = years

	for year := 1; year <= years; year++ {
		if year == 1 {
			result[1] += 15
			result[2] += 15
		} else if year == 2 {
			result[1] += 9
			result[2] += 9
		} else {
			result[1] += 4
			result[2] += 5
		}
	}
	return result
}

func main() {
	fmt.Println(CalculateYears(5))
}
