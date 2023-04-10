package convert

func ArrayToMap[T any, V comparable](array []T, key func(T) V) map[V]T {
	result := make(map[V]T, len(array))

	for _, element := range array {
		result[key(element)] = element
	}

	return result
}

func MapToMapArray[T any, V comparable](currentMap map[V]T, key func(T) V) map[V][]T {
	result := make(map[V][]T)

	for _, element := range currentMap {
		currentKey := key(element)
		result[currentKey] = append(result[currentKey], element)
	}

	return result
}
