package main

import (
	"fmt"
	"os"
	"strconv"
)

func intToRoman(num int) (string, string) {
	if num <= 0 || num >= 4000 {
		return "", "error :  cannot convert to roman digit"
	}
	romanSymbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	romanValues := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	var romancalculation string
	var romanresult string
	for i := 0; i < len(romanSymbols); i++ {
		for num >= romanValues[i] {

			romancalculation += romanSymbols[i] + "+"
			romanresult += romanSymbols[i]
			num -= romanValues[i]

		}
	}
	if len(romancalculation) > 0 {
		romancalculation = romancalculation[:len(romancalculation)-1]
	}
	return romancalculation, romanresult
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Message Error")
		os.Exit(1)
	}
	numStr := os.Args[1]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Message error cannort convert to romain digit")
		os.Exit(1)
	}
	calculation, 
	result := intToRoman(num)
	if calculation == "" {
		fmt.Println(result)
	} else {
		fmt.Println(calculation)
		fmt.Println(result)
	}
}
