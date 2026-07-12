package player

import "strings"

type growthRateFunctionId int

const (
	slow growthRateFunctionId = iota
	medium
	fast
	mediumSlow
	erratic
	fluctuating
)

var growthRateFunctions = [6]func(level int) (amount int){
	slowGrowthRate,
	mediumGrowthRate,
	fastGrowthRate,
	mediumSlowGrowthRate,
	erraticGrowthRate,
	fluctuatingGrowthRate,
}

func getGrowthRateFunction(name string) func(level int) (amount int) {
	name = strings.ToLower(name)
	idx := -1
	switch name {
	case "slow":
		idx = int(slow)
	case "medium":
		idx = int(medium)
	case "fast":
		idx = int(fast)
	case "medium-slow":
		idx = int(mediumSlow)
	case "slow-then-very-fast":
		idx = int(erratic)
	case "fast-then-very-slow":
		idx = int(fluctuating)
	default:
		idx = int(medium)
	}
	return growthRateFunctions[idx]
}

func slowGrowthRate(level int) (amount int) {
	return 5 * intPow(level, 3) / 4
}

func mediumGrowthRate(level int) (amount int) {
	return intPow(level, 3)
}

func fastGrowthRate(level int) (amount int) {
	return 4 * intPow(level, 3) / 5
}

func mediumSlowGrowthRate(level int) (amount int) {
	return 6*intPow(level, 3)/5 + 15*intPow(level, 2) + 100*level - 140
}

func erraticGrowthRate(level int) (amount int) {
	if level <= 50 {
		return (intPow(level, 3) * (100 - level)) / 50
	} else if level <= 68 {
		return (intPow(level, 3) * (150 - level)) / 100
	} else if level <= 98 {
		return (intPow(level, 3) * (1274 + intPow((level%3), 2) - 9*(level%3) - 20*(level/3))) / 1000
	} else {
		return (intPow(level, 3) * (160 - level)) / 100
	}
}

func fluctuatingGrowthRate(level int) (amount int) {
	if level < 15 {
		return (intPow(level, 3) * ((level+1)/3 + 24)) / 50
	} else if level < 36 {
		return (intPow(level, 3) * (level + 14)) / 50
	} else {
		return (intPow(level, 3) * ((level / 2) + 32)) / 50
	}
}

// TODO(emi): write a test for this
func intPow(num, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return num
	}

	result := num
	for i := 2; i <= n; i++ {
		result *= num
	}
	return result
}
