package array_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"penahub.gitlab.yandexcloud.net/pena-services/accruals-service/pkg/array"
)

func TestIndexOf(t *testing.T) {
	t.Run("Find index of element in array (success)", func(t *testing.T) {
		callback := func(element string, index int, array []string) bool {
			assert.Equal(t, array, []string{"test1", "test2"})

			return element == "test2"
		}

		assert.Equal(t, 1, array.IndexOf([]string{"test1", "test2"}, callback))
	})

	t.Run("Find index of element in array (failure)", func(t *testing.T) {
		callback := func(element string, index int, array []string) bool {
			assert.Equal(t, array, []string{"test1", "test2"})

			return element == "test"
		}

		assert.Equal(t, 0, array.IndexOf([]string{"test1", "test2"}, callback))
	})
}
