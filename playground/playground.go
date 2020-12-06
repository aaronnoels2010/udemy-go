package main

import "fmt"

type hotdog int
var z hotdog

func main() {
	z = 43
	fmt.Println(z)
	fmt.Printf("%T", z)

}

func foo() {
	fmt.Println("I am just a foo")
}
