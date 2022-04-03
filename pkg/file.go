package pkg

import (
	"os"
	"strings"
	"sync"
)

func AddStringToFile(filename string, data string, mutex *sync.RWMutex) error {
	mutex.Lock()

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	_, err = file.WriteString(data + "\n")
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	mutex.Unlock()

	return nil
}

func GetAllStringsFromFile(filename string, mutex *sync.RWMutex) ([]string, error) {
	mutex.Lock()

	file, err := os.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}
	strArr := strings.Split(string(file), "\n")

	mutex.Unlock()

	return strArr, err
}
