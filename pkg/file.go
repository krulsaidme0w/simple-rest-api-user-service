package pkg

import (
	"errors"
	"os"
	"strings"
	"sync"
)

var FileAlreadyExists = errors.New("FileAlreadyExists")

func CreateFileWithData(path string, filename string, data string, mutex *sync.RWMutex) error {
	if _, err := os.Stat(path + "/" + filename); errors.Is(err, os.ErrNotExist) {
		mutex.Lock()

		file, err := os.OpenFile(path+"/"+filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}

		_, err = file.WriteString(data)
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

	return FileAlreadyExists
}

func WriteStringsToFile(path string, stringArr []string, mutex *sync.RWMutex) error {
	mutex.Lock()

	if err := os.Truncate(path, 0); err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_APPEND, 0664)
	if err != nil {
		return err
	}

	sep := "\n"
	for index, str := range stringArr {
		if index == len(stringArr)-1 {
			sep = ""
		}

		if _, err = file.WriteString(str + sep); err != nil {
			return err
		}
	}

	err = file.Close()
	if err != nil {
		return err
	}

	mutex.Unlock()

	return nil
}

func GetFirstStringFromFile(path string, mutex *sync.RWMutex) (string, error) {
	mutex.Lock()

	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	mutex.Unlock()

	return string(file), nil
}

func GetAllStringsFromFile(path string, mutex *sync.RWMutex) ([]string, error) {
	mutex.Lock()

	file, err := os.ReadFile(path)
	if err != nil {
		return []string{}, err
	}
	strArr := strings.Split(string(file), "\n")

	mutex.Unlock()

	return strArr, nil
}
