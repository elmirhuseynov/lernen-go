package main

import "fmt"

var romanNumeralMap = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

func intToRoman(num int) string {
	result := ""
	for value, symbol := range romanNumeralMap {
		for num >= value {
			result += symbol
			num -= value
		}
	}
	return result
}

func main() {
	num := 1234
	fmt.Println(intToRoman(num))
}
