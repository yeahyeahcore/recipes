package array

func IndexOf[T any](array []T, callback func(T, int, []T) bool) int {
	for index, element := range array {
		if callback(element, index, array) {
			return index
		}
	}

	return 0
}
