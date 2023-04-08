package array_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"penahub.gitlab.yandexcloud.net/pena-services/accruals-service/pkg/array"
)

func TestFind(t *testing.T) {
	t.Run("Find element in array (success)", func(t *testing.T) {
		result := "test1"
		callback := func(element string, index int, array []string) bool {
			assert.Equal(t, array, []string{"test1", "test2"})

			return element == result
		}

		assert.Equal(t, &result, array.Find([]string{"test1", "test2"}, callback))
	})

	t.Run("Find element in array (failure)", func(t *testing.T) {
		result := "test"
		callback := func(element string, index int, array []string) bool {
			assert.Equal(t, array, []string{"test1", "test2"})

			return element == result
		}

		assert.Nil(t, array.Find([]string{"test1", "test2"}, callback))
	})
}
