package string_calculator

import (
	"strconv"
	"strings"
)

func Add(numbers string) int {
	if len(numbers) == 0 {
		return 0
	}

	result := 0
	numbers = strings.Replace(numbers,"\n",",",1)
	for _, v := range strings.Split(numbers, ",") {
		n, _ := strconv.Atoi(v)
		result += n
	}
	return result
}
