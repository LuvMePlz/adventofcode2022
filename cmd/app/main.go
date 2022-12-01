package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filePath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	filePath += "\\data\\input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		fmt.Println(filePath)
	}
	defer file.Close()

	var max int
	var value int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			max = updateMax(max, value)
			value = 0
			continue
		}

		var temp, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}
		value += temp
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(max)
}

func updateMax(max int, value int) int {
	if max < value {
		max = value
		fmt.Println(max)
	}
	return max
}
