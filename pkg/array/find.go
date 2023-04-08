package array

func Find[T any](array []T, callback func(T, int, []T) bool) *T {
	for index, element := range array {
		if callback(element, index, array) {
			return &element
		}
	}

	return nil
}
