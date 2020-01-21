package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) setX(x float64) {
	c.x = x
}

func main() {
	c := new(Circle)
	d := Circle{x: 0, y: 0, r: 5}
	e := &Circle{x: 0, y: 0, r: 0}
	f := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f == d)
	fmt.Println(c.area())
	fmt.Println(d.area())

	c.setX(5)
	d.setX(4)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(totalArea(c, e))
}

func totalArea(shapes ...Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.area()
	}

	return total
}
