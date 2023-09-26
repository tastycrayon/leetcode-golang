package main

func multiply(num1 string, num2 string) string {
	length1, length2 := len(num1), len(num2)
	const zero byte = 48

	if (length1 == 1 && num1[0] == zero) || (length2 == 1 && num2[0] == zero) {
		return "0"
	}

	maxLength := length1 + length2
	result := make([]byte, maxLength)

	for i := length2 - 1; i >= 0; i-- {
		var digit2 byte = num2[i] - zero
		for j := length1 - 1; j >= 0; j-- {
			var digit1 byte = num1[j] - zero
			result[i+j+1] += digit1 * digit2  // last digit
			result[i+j] += result[i+j+1] / 10 // has carry
			result[i+j+1] %= 10               // single digit
		}
	}
	output := make([]byte, 0, maxLength)

	var checkForZero = true
	for _, val := range result {
		if val != 0 {
			checkForZero = false
		}
		if checkForZero {
			continue
		}
		output = append(output, val+zero)
	}
	return string(output)
}
