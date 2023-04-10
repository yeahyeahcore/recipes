package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeahyeahcore/recipes/pkg/utils"
)

func TestMergeMaps(t *testing.T) {
	t.Run("Соединение двух пустых map", func(t *testing.T) {
		assert.Equal(t, map[string]interface{}{}, utils.MergeMaps(map[string]interface{}{}, map[string]interface{}{}))
	})

	t.Run("Соединение несколько пустых map", func(t *testing.T) {
		assert.Equal(t, map[string]interface{}{}, utils.MergeMaps(map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}))
	})

	t.Run("Соединение заполненных map", func(t *testing.T) {
		assert.Equal(t,
			map[string]interface{}{
				"test1": "1",
				"test2": "1",
				"test3": "1",
				"test4": "1",
			},
			utils.MergeMaps(
				map[string]interface{}{"test1": "1"},
				map[string]interface{}{"test2": "1"},
				map[string]interface{}{"test3": "1", "test4": "1"},
			),
		)
	})

	t.Run("Соединение заполненных map с одинаковыми ключами", func(t *testing.T) {
		assert.Equal(t,
			map[string]interface{}{
				"test1": "1",
				"test2": "1",
				"test3": "1",
				"test4": "1",
			},
			utils.MergeMaps(
				map[string]interface{}{"test1": "1"},
				map[string]interface{}{"test2": "1", "test1": "1"},
				map[string]interface{}{"test3": "1", "test4": "1"},
			),
		)
	})
}
