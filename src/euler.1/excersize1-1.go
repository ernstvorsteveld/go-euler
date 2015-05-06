package main

import (
	"euler.1/mulitpliers"
	"fmt"
)

func main() {
	var multiplies = Multiplies{[]int{3, 5}, 1000}
	var result = findMultipliers(multiplies)
	fmt.Println("result: ", result)
}
