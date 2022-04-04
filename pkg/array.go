package pkg

import "strings"

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
