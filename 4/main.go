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
	println(a)
	println(b)
	println(c)
	println(d)
	println(f)
	fmt.Printf("O tipo de C é %T ", f)

}

