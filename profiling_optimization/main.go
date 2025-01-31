package main

import "math/rand"

func heavyComputation() {
	var result int
	for i := 0; i < 1000000000; i++ {
		result += rand.Intn(100)
	}
}
