package main

import "fmt"

func main() {
	var b int
	b = 7
	fmt.Println(half(b))
}

func half(a int) (b int, x bool) {
	b = a / 2

	if a%2 == 1 {
		x = false
	} else {
		x = true
	}
	return b, x
}
