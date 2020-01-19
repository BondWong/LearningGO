package main

import "fmt"

var g = "Gloabl String"

func f() {
	fmt.Println(g)
}

func main() {
	var x string
	x = "Hello"
	x = x + " World"

	var y string
	y = "Hello"
	y += " World"

	z := "hi"
	a := 1

	const b = 2

	var (
		c = "hey"
		d = "hi"
		e = "hello"
	)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(y == x)
	fmt.Println(z)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c, d, e)
	f()

	fahrenheitToCelsius()
	feetToMeters()
}

func fahrenheitToCelsius() {
	fmt.Println("Input a fahrentheit degree")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)
	fmt.Println((fahrenheit - 32) * 5 / 9)
}

func feetToMeters() {
	fmt.Println("Input feet")
	var feet float64
	fmt.Scanf("%f", &feet)
	fmt.Println(0.3048 * feet)
}
