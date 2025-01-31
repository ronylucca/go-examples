package main

import "fmt"

func main() {
	var minhaVar interface{} = "Texto"
	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o valor de Ok é %v\n", res, ok)

	res2 := minhaVar.(int)
	fmt.Printf("O valor de res2 é %v", res2)
}
