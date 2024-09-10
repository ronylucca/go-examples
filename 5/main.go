package main

import fmt "fmt"
type ID int
var (
	a bool = true	
	b int = 1
	c float64 = 1.0
	d string = "Hello, World!"
	f ID = 1
)

func main() {
	var meuArray [4]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30
	meuArray[3] = 40
	fmt.Println(meuArray)
	fmt.Println(meuArray[0])
	fmt.Println(meuArray[(len(meuArray)-1)])

	for i, v:= range meuArray {
		fmt.Printf("O valor do indice e %d . E o valor é %d\n" , i , v)
	}

}

