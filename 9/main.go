package main

import fmt "fmt"

func main() {

	total := soma(10, 11, 2332, 2)
	fmt.Printf("A soma de %d\n", total)
}

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
