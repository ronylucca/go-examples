package main

func main() {
	var1 := 10
	var2 := 20
	print(soma(&var1, &var2))
}

func soma(a, b *int) int {
	*a = 30
	*b = 70
	return (*a + *b)
}
