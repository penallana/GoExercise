package main

import "fmt"

func main() {

	halfi := func(b int) (int, bool) {
		return b / 2, b%2 == 1 //true if odd
		return b / 2, b%2 == 0 //true if even

	}
	fmt.Println(halfi(7))
}
