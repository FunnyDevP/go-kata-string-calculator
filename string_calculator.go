package string_calculator

import (
	"regexp"
	"strconv"
)

func Add(numbers string) int {
	if len(numbers) == 0 {
		return 0
	}

	result := 0
	re := regexp.MustCompile(`(?m)[0-9]+`)
	numbersArr := re.FindAllString(numbers,-1)
	for _, v := range numbersArr {
		n, _ := strconv.Atoi(v)
		result += n
	}
	return result
}
