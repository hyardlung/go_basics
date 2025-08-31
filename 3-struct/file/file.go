package file

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func ReadFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	ext := filepath.Ext(fileName)
	isValidJSON := json.Valid(data)
	if err != nil {
		return nil, err
	}
	if ext != ".json" || !isValidJSON {
		err = fmt.Errorf("can't read %s file type. Only '.json' files available", ext)
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