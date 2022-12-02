package inputReader

import (
	"fmt"
	"os"
)

func OpenInputFileReader(addedFilePath string) (*os.File, error) {
	filePath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	filePath += addedFilePath

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		fmt.Println(filePath)
		return nil, err
	}
	return file, nil
}
