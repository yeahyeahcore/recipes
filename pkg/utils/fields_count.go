package utils

import (
	"reflect"
)

func GetFilledFieldsCount(object interface{}) int {
	filledFieldsCount := int(0)

	if object == nil {
		return 0
	}

	value := reflect.ValueOf(object)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	fieldsCount := value.NumField()
	if fieldsCount < 1 {
		return 0
	}

	for index := 0; index < fieldsCount; index++ {
		field := value.Field(index)

		if field.IsZero() {
			continue
		}

		filledFieldsCount++
	}

	return filledFieldsCount
}
