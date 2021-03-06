package multipliers

type Multiplies struct {
	multipliers []int
	max         int
}

func isMultiplier(value int, max int) bool {
	return max%value == 0
}

func findMultipliers(multipliers Multiplies) int {
	var sum = 0
	for i := 1; i < multipliers.max; i++ {
		var isOk = false
		for j := range multipliers.multipliers {
			isOk = isOk || isMultiplier(multipliers.multipliers[j], i)
		}
		if isOk {
			sum += i
		}
	}
	return sum
}
