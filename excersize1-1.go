package main

import (
	"fmt"
	"mulitpliers"
)

func main() {
	var multiplies = Multiplies{[]int{3, 5}, 1000}
	var result = findMultipliers(multiplies)
	fmt.Println("result: ", result)
}
