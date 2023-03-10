package main

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	length1, length2 := len(num1), len(num2)

	if length1 < length2 {
		num1, num2 = num2, num1
		length1, length2 = len(num1), len(num2)
	}
	finalResult := ""
	carry := 0

	for i := length2 - 1; i >= 0; i-- {
		result := 0
		remainder := 0
		number2, _ := strconv.Atoi(string(num2[i]))

		multiplier := 0
		for j := length1 - 1; j >= 0; j-- {
			jPlace := (length1 - 1 - j)
			if jPlace == 0 {
				multiplier = 1
			} else {
				multiplier *= 10
			}

			number1, _ := strconv.Atoi(string(num1[j]))

			tempResult := number2 * number1
			tempResult += remainder // add remainder

			oneResult := tempResult % 10
			remainder = tempResult / 10

			result = oneResult*multiplier + result

		}
		if remainder != 0 {
			carry += remainder * multiplier * 10
		}
		temp := result + carry
		result = temp % 10
		carry = temp - result
		carry = carry / 10

		finalResult = strconv.Itoa(result) + finalResult
	}

	if carry != 0 {
		return strconv.Itoa(carry) + finalResult
	}
	return finalResult
}

func main() {
	fmt.Println(multiply("925101087184894", "3896737933784656127"))
	//3604876499018802870538090258945538
}
