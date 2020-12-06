package main

import "fmt"

func main() {
	x := "Dit is de enige uitweg die ik zie"
	fmt.Println([]byte(x))
	for index, value := range x {
		fmt.Printf("index %d, %d and %#x\n", index, value, value)
	}
}

func exercise01_02() {
	x := 42
	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%x\n", x)
}

func exercise02_02() {
	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 41)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 41)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

func exercise03_02() {
	const(
		x = 42
		y int32 = 5
	)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}

func exercise04_02() {
	x := 42
	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%x\n", x)
	a := x << 1
	fmt.Printf("%d\n", a)
	fmt.Printf("%b\n", a)
	fmt.Printf("%#X\n", a)
}

func exercise05_02() {
	x := `Dit is
	een random string
	die letterlijk genomen wordt.`
	fmt.Println(x)
}

func exercise06_02() {
	const (
		x = 2020
		a = x + iota
		b = x + iota
		c = x + iota
		d = x + iota
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}