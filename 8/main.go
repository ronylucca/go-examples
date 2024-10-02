package main

import fmt "fmt"

func main() {

	total, isBiggerThanTwenty := soma(10, 11)
	fmt.Printf("A soma de 10 + 11 é %d. Maior que 20? %t\n", total, isBiggerThanTwenty)
}

func soma(a int, b int) (int, bool) {
	if (a + b) > 20 {
		return a + b, true
	}
	return a + b, false
}
