package main

import "fmt"

func main() {
	a := 3
	b := 5
	s := add(a, b)

	m := multiply(a, b)

	s1, m1 := math(a, b, add, multiply)
	fmt.Println(s, m, s1, m1)

}

//check
func add(x, y int) int {
	z := x + y
	return z
}
func multiply(x, y int) int {
	z := x * y
	return z
}
func math(a, b int, f1 func(x, y int) int, f2 func(a, b int) int) (int, int) {
	return f1(a, b), f2(a, b)
}
