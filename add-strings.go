package main

// Input: num1 = "91", num2 = "123"
// Output: "214"
func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	var offset int = len(num1) - len(num2)
	const zero byte = '0' // 48
	result := make([]byte, len(num1)+1)

	for i := len(num1) - 1; i >= 0; i-- {
		var d1 byte = num1[i] - zero // digit is in int, but type is byte
		var d2 byte = 0
		if i-offset >= 0 {
			d2 = num2[i-offset] - zero // if not in-bound, use zero otherwise proper value
		}
		result[i+1] += d1 + d2       // last index
		result[i] = result[i+1] / 10 // carry (one index before)
		result[i+1] %= 10            // make it single digit
		result[i+1] += zero          // make ascii
	}
	if result[0] == 0 {
		result = result[1:]
	} else {
		result[0] += zero
	}

	return string(result)
}
