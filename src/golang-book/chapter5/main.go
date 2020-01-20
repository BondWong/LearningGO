package main

import "fmt"

func main() {
	var total float64
	x := [5]float64{1, 2, 3, 4, 10}
	for _, num := range x {
		total += num
	}
	fmt.Println(total)

	y := []float64{1, 2, 3, 4, 5, 6, 8}
	fmt.Println(y[1:])

	z := make(map[string][]int)
	z["first"] = []int{10, 20}
	z["second"] = []int{10, 20, 30}
	delete(z, "second")
	fmt.Println(z["second"]) // return zero value

	if arr, ok := z["first"]; ok {
		fmt.Println(arr)
	}
}
