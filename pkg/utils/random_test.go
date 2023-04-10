package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yeahyeahcore/recipes/pkg/utils"
)

func TestGetRandomString(t *testing.T) {
	t.Run("Создание рандомной строки с размерностью 0", func(t *testing.T) {
		randomString := utils.GetRandomString(0)

		assert.Equal(t, 0, len(randomString))
	})

	t.Run("Создание рандомной строки", func(t *testing.T) {
		randomString := utils.GetRandomString(8)

		assert.Equal(t, 8, len(randomString))
	})
}
