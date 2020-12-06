package main

import "fmt"

type Exercise04 int

func main() {
	exercise05()
}

func exercise01() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func exercise02() {
	var x int
	var y string
	var z bool
	fmt.Printf("%T : %v \n", x, x)
	fmt.Printf("%T : %v \n", y, y)
	fmt.Printf("%T : %v \n", z, z)
}

func exercise03() {
	var x int
	var y string
	var z bool
	x = 42
	y = "James Bond"
	z = true
	result := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(result)
}

func exercise04() {
	var x Exercise04
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 43
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
}

func exercise05() {
	var x Exercise04
	x = 43
	y := int(x)
	fmt.Printf("value: %v, type: %T\n", y, y)
}

func exercise06() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
