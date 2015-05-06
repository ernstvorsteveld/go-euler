package main

import (
	"fmt"
)

// Find the sum of all the multiples of 3 or 5 below 1000.

type Multiplies struct {
	multipliers []int
	max         int
}

func findMultipliesForInt(base int, multiplier int, max int) []int {
	if multiplier < max {
		var x []int = make([]int, 0)
		if max/multiplier > 0 {
			x = append(x, multiplier)
		}
		var y []int = findMultipliesForInt(base, base+multiplier, max)

		return append(x, y...)
	}
	return []int{}
}

func findMultiplies(base int, multipliers int, max int) []int {
	var x []int = make([]int, 0)
	var value = findMultipliesForInt(base, multipliers, max)
	x = append(x, value...)
	return x
}

func calculateSum(values []int) int {
	fmt.Println(values)
	var sum = 0
	for i := range values {
		sum += values[i]
	}
	fmt.Println(sum)
	return sum
}

func main() {
	var m = Multiplies{multipliers: []int{3, 5}, max: 1000}
	fmt.Println(10 / 3)
	fmt.Println(m)

	var sum int = 0
	for i := range m.multipliers {
		var base = m.multipliers[i]
		var multiplier = m.multipliers[i]
		var max = m.max
		var value = findMultipliesForInt(base, multiplier, max)

		sum = sum + calculateSum(value)
	}

	fmt.Println(sum)
}
