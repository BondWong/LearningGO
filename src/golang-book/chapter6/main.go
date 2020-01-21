package main

import "fmt"

func main() {
	d := func() {
		error := recover()
		fmt.Println(error)
	}
	defer d()
	x := []float64{1, 2, 3, 4, 4}
	fmt.Println(average(x))
	fmt.Println(add(x...))

	panic("I am panic")
}

func average(arr []float64) (ave float64) {
	ave = 0.0
	if len(arr) == 0 {
		return
	}

	for _, num := range arr {
		ave += num
	}

	ave = ave / float64(len(arr))
	return
}

func add(nums ...float64) float64 {
	sum := 0.0
	for _, num := range nums {
		sum += num
	}

	return sum
}
