package string_calculator

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func Add(numbers string) (int,error) {
	if len(numbers) == 0 {
		return 0,nil
	}

	result := 0
	re := regexp.MustCompile(`(?m)[0-9]+|-\d+`)
	numbersArr := re.FindAllString(numbers,-1)

	negativeVal := make([]string,0)
	for _, v := range numbersArr {
		n, _ := strconv.Atoi(v)

		if n < 0 {
			negativeVal = append(negativeVal,v)
		}
		result += n
	}

	if len(negativeVal) > 0 {
		return 0, errors.New("error: negatives not allowed: " + strings.Join(negativeVal," "))
	}
	return result,nil
}
