package daytwo

import (
	"adventofcode/internal/inputReader"
	"strings"

	"bufio"
	"fmt"
)

func TaskOne() int {
	file, err := inputReader.OpenInputFileReader("\\data\\input2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var equalsMap = map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	var winnerMap = map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}

	var pointsMap = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	scanner := bufio.NewScanner(file)
	var result = 0
	for scanner.Scan() {
		// scanner.Scan()
		var symbols = strings.Split(scanner.Text(), " ")
		result += pointsMap[symbols[1]]

		// fmt.Println(symbols[0], " : ", symbols[1])
		// fmt.Println("symbols[0] == symbols[1] = ", symbols[0] == symbols[1])
		if symbols[1] == equalsMap[symbols[0]] {
			result += 3
			continue
		}

		// fmt.Println("symbols[1] == winnerMap[symbols[0]] = ", symbols[1] == winnerMap[symbols[0]])
		if symbols[1] == winnerMap[symbols[0]] {
			result += 6
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return result
}

func TaskTwo() int {
	file, err := inputReader.OpenInputFileReader("\\data\\input2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var symbolPointsMap = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	var pointsMap = map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	var symbolPoints = [3]int{1, 2, 3}
	scanner := bufio.NewScanner(file)
	var result = 0
	for scanner.Scan() {
		var symbols = strings.Split(scanner.Text(), " ")
		result += pointsMap[symbols[1]]

		switch symbols[1] {
		case "X":
			var index = symbolPointsMap[symbols[0]] - 1
			var length = len(symbolPoints)
			result += symbolPoints[(index+length-1)%length]
		case "Y":
			result += symbolPointsMap[symbols[0]]
		case "Z":
			var index = symbolPointsMap[symbols[0]] - 1
			var length = len(symbolPoints)
			result += symbolPoints[(index+1)%length]
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return result
}
