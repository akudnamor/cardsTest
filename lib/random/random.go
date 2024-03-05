package random

import (
	"math/rand"
)

func RandomInt() int {
	return rand.Intn(100)
}

func RandomString() string {
	var result string
	var source = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	for i := 0; i < 3; i++ {
		r := rand.Intn(len(source))
		result += string(source[r])
	}
	return result
}

func RandomFloat64() float64 {
	return rand.Float64() * 10
}
