package dayfour

import (
	"adventofcode/internal/inputReader"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func TaskOne() int {
	file, err := inputReader.OpenInputFileReader("\\data\\input4.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result = 0
	for scanner.Scan() {
		var parts = strings.Split(scanner.Text(), ",")
		firstStart, firstEnd := getStartEnd(parts[0])
		secondStart, secondEnd := getStartEnd(parts[1])

		if (firstStart <= secondStart && firstEnd >= secondEnd) ||
			(secondStart <= firstStart && secondEnd >= firstEnd) {
			result++
			continue
		}

	}
	return result
}

func getStartEnd(scope string) (int, int) {
	var limits = strings.Split(scope, "-")
	start, err := strconv.Atoi(limits[0])
	if err != nil {
		fmt.Println(err)
	}
	end, err := strconv.Atoi(limits[1])
	if err != nil {
		fmt.Println(err)
	}
	return start, end
}

func TaskTwo() int {
	file, err := inputReader.OpenInputFileReader("\\data\\input4.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result = 0
	for scanner.Scan() {
		var parts = strings.Split(scanner.Text(), ",")
		firstStart, firstEnd := getStartEnd(parts[0])
		secondStart, secondEnd := getStartEnd(parts[1])

		if firstStart <= secondEnd && firstEnd >= secondStart {
			result++
			continue
		}

	}
	return result
}
