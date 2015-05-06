package main

import (
	"fmt"
)

// Find the sum of all the multiples of 3 or 5 below 1000.

type Multiplies struct {
	multipliers []int
	max         int
}

type IntSet struct {
	set map[int]bool
}

func NewIntSet() *IntSet {
	return &IntSet{make(map[int]bool)}
}

func (set *IntSet) Add(i int) bool {
	_, found := set.set[i]
	set.set[i] = true
	return !found //False if it existed already
}

func findMultipliesForInt(set *IntSet, base int, multiplier int, max int) {
	if multiplier < max {
		if max/multiplier > 0 {
			set.Add(multiplier)
		}
		findMultipliesForInt(set, base, base+multiplier, max)
	}
}

func calculateSum(set *IntSet, max int) int {
	fmt.Println(set)
	var sum = 0
	for k, v := range set.set {
		if v {
			sum += k
		}
	}
	fmt.Println(sum)
	return sum
}

func main() {
	var m = Multiplies{multipliers: []int{3, 5}, max: 1000}
	set := NewIntSet()
	var max = m.max
	for i := range m.multipliers {
		var base = m.multipliers[i]
		var multiplier = m.multipliers[i]
		findMultipliesForInt(set, base, multiplier, max)
	}
	fmt.Println("som:", calculateSum(set, max))
}
