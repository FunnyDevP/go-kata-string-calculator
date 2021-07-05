package string_calculator

import "strconv"

func Add(numbers string) int {
	if len(numbers) == 0 {
		return 0
	}

	n, _ := strconv.Atoi(numbers)
	return n
}
