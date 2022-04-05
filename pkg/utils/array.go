package utils

import (
	"errors"
	"strings"
)

var CannotFindInArray = errors.New("CannotFindInArray")

func search(array []string, search string) (int, error) {
	low := 0
	high := len(array) - 1

	for low <= high {
		median := (low + high) / 2
		elem := strings.Split(array[median], " ")[1]

		if elem == search {
			return median, nil
		}

		if elem < search {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	return -1, CannotFindInArray
}

func binarySearch(array []string, search string) int {
	if len(array) == 0 {
		return -1
	}

	left := 0
	right := len(array) - 1

	for left <= right {
		mid := (left + right) / 2
		midElem := strings.Split(array[mid], " ")[1]

		if midElem < search {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func FindInSortedArray(array []string, data string) (string, error) {
	index, err := search(array, data)
	if err != nil {
		return "", CannotFindInArray
	}

	return strings.Split(array[index], " ")[0], nil
}

func InsertToArray(array []string, data string) []string {
	index := binarySearch(array, strings.Split(data, " ")[1])
	if index < 0 {
		return append([]string{data}, array...)
	}

	if index > len(array)-1 {
		return append(array, data)
	}

	array = append(array, "")
	copy(array[index+1:], array[index:])
	array[index] = data
	return array
}
