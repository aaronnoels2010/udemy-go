package main

import "fmt"

func main() {
	exercise10_03()
}

func exercise01_03() {
	for i := 1; i < 10000; i++ {
		fmt.Println(i)
	}
}

func exercise02_03() {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for index, runeValue := range letters {
		fmt.Println(index)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", runeValue)
		}
	}
}

func exercise03_03() {
	for year := 1930; year <= 2020 ; year++ {
		if year >= 1999 && year <= 2020 {
			fmt.Println(year)
		}
	}
}

func exercise04_03() {
	year := 1930
	for {

		if year >= 1999 && year <= 2020 {
			fmt.Println(year)
		}

		if year == 2020 {
			break
		}

		year++
	}
}

func exercise05_03() {
	for i := 11; i < 100; i++ {
		fmt.Printf("%v: %v\n", i, i % 4)
	}
}

func exercise06_03() {
	fmt.Println("Exercise06_03")
	if 2 == 2 {
		fmt.Println("\tif 2 == 2")
	}
}

func exercise07_03() {
	x := 42

	if x == 40 {
		fmt.Println("The number was 40.")
	} else if x == 41 {
		fmt.Println("The number was 41.")
	} else {
		fmt.Printf("The number wasn't 40 or 41 but: %v\n", x)
	}
}

func exercise08_03() {
	switch {
	case false:
		fmt.Println("The expression was false.")
	case true:
		fmt.Println("The expression is true.")
	}
}

func exercise09_03() {
	switch "favSport" {
	case "favSport":
		fmt.Println("The expression was favSport")
	default:
		fmt.Println("The default expression")
	}
}

func exercise10_03() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}