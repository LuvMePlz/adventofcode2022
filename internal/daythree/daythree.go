package daythree

import (
	"adventofcode/internal/inputReader"
	"bufio"
	"fmt"
	"unicode"
)

func TaskOne() int {
	file, err := inputReader.OpenInputFileReader("\\data\\input3.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var lowerCaseShift = 96
	var upperCaseShift = 38
	scanner := bufio.NewScanner(file)
	var result = 0
	for scanner.Scan() {
		firstHalfMap, secondHalfMap := fillSymbolsMaps(scanner.Text(), len(scanner.Text()))

		for _, v := range scanner.Text() {
			if firstHalfMap[byte(v)] && secondHalfMap[byte(v)] {
				if unicode.IsLower(v) {
					result += int(v) - lowerCaseShift
					break
				} else {
					result += int(v) - upperCaseShift // fmt.Println(v, "U")
					break
				}
			}
		}
	}
	return result
}

func fillSymbolsMaps(symbols string, length int) (map[byte]bool, map[byte]bool) {
	var firstHalf = make(map[byte]bool, length/2)
	var secondHalf = make(map[byte]bool, length/2)
	var tempMap = firstHalf
	for i, v := range symbols {
		if i == length/2 {
			tempMap = secondHalf
		}
		tempMap[byte(v)] = true
	}
	return firstHalf, secondHalf
}

func TaskTwo() int {
	file, err := inputReader.OpenInputFileReader("\\data\\input3.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var i = 0
	var symbols = make(map[byte]int, 20)
	var result = 0
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		for _, v := range scanner.Text() {
			if symbols[byte(v)] == i {

				if i == 2 {
					result = calculateResult(result, v)
					// fmt.Println(result, v)
					break
				}
				symbols[byte(v)]++
			}
		}

		// fmt.Println(symbols)

		if i == 2 {
			symbols = make(map[byte]int, 20)
			i = 0
			continue
			// break
		}
		i++

	}
	return result
}

func calculateResult(result int, v rune) int {
	var lowerCaseShift = 96
	var upperCaseShift = 38

	if unicode.IsLower(v) {
		result += int(v) - lowerCaseShift
	} else {
		result += int(v) - upperCaseShift // fmt.Println(v, "U")
	}
	return result
}
