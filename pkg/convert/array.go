package convert

import "strconv"

func StringArrayToIntSlice(stringArray []string) []int {
	intArray := make([]int, 0, len(stringArray))

	for _, stringValue := range stringArray {
		intValue, err := strconv.Atoi(stringValue)
		if err != nil {
			continue
		}

		intArray = append(intArray, intValue)
	}
	return intArray
}

func MapToArray[T any, V comparable](currentMap map[V]T) []T {
	result := make([]T, len(currentMap))
	index := 0

	for _, element := range currentMap {
		result[index] = element
		index++
	}

	return result
}
