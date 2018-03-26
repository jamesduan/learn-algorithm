package main

import "fmt"

/**
 * @param number: A 3-digit number.
 * @return: Reversed number.
 */
func reverseInteger(number int) int {
	// write your code here
	hundreds := number / 100
	tens := (number / 10) % 10
	units := number % 10
	result := units*100 + tens*10 + hundreds
	return result
}

func lowerCaseToUppercase(character byte) byte {
	step := 97 - 65
	return byte(int(character) - step)
}

func main() {
	// reverseInteger(123)
	result := lowerCaseToUppercase('b')
	fmt.Println(string(result))
}
