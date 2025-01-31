package main

import "fmt"

func main() {

	var x interface{} = 10
	var y interface{} = "hello"

	showType(x)
	showType(y)
}

func showType(i interface{}) {
	fmt.Printf("O tipo da variavel é %T e o valor é %v\n", i, i)
}
