package main

import "fmt"

type hotdog int

var x hotdog

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	y := 30
	x = hotdog(y)
	fmt.Println(x)

}
