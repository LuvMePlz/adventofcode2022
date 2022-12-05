package dayfour

import (
	"adventofcode/internal/inputReader"
	"bufio"
	"fmt"
)

func TaskOne() {
	file, err := inputReader.OpenInputFileReader("\\data\\input3.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

	}
}
