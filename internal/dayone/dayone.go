package dayone

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TaskOne() int {
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

	var max = 0
	var value = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			if max < value {
				max = value
			}
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

	fmt.Println("max=", max)
	return max
}

func TaskTwo() int {
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

	var max = make([]int, 3)
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

	var result = 0
	for _, v := range max {
		result += v
	}
	fmt.Println("max=", max)
	fmt.Println("result=", result)
	return result
}

func updateMax(max []int, value int) []int {
	for i, v := range max {
		if v < value {
			max = append(max[0:i+1], max[i:2]...)
			max[i] = value
			//fmt.Println(max)
			break
		}
	}
	return max
}
