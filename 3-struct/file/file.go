package file

import (
	"fmt"
	"os"
)

func ReadFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteFile(content []byte, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
}