package main

func main() {
	a := 10
	ponteiroA := &a
	*ponteiroA = 20

	println(a)
}
