package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yeahyeahcore/recipes/pkg/utils"
)

func TestGetFilledFieldsCount(t *testing.T) {
	name := "name"
	age := uint(10)

	t.Run("Получение количества заполненных полей структуры с пустой структурой", func(t *testing.T) {
		assert.Equal(t, 0, utils.GetFilledFieldsCount(struct{}{}))
	})

	t.Run("Получение количества заполненных полей структуры с nil", func(t *testing.T) {
		assert.Equal(t, 0, utils.GetFilledFieldsCount(nil))
	})

	t.Run("Получение количества заполненных полей структуры с указателем на пустую структуру", func(t *testing.T) {
		assert.Equal(t, 0, utils.GetFilledFieldsCount(&struct{}{}))
	})

	t.Run("Получение количества заполненных полей структуры с заполненной структурой", func(t *testing.T) {
		assert.Equal(t, 1, utils.GetFilledFieldsCount(struct {
			name *string
			age  *uint
		}{
			name: &name,
		}))
	})

	t.Run("Получение количества заполненных полей структуры с указателем на заполненную структурой", func(t *testing.T) {
		assert.Equal(t, 2, utils.GetFilledFieldsCount(&struct {
			name *string
			age  *uint
		}{
			name: &name,
			age:  &age,
		}))
	})

	t.Run("Получение количества заполненных полей структуры с заполненную структурой не опциональных значений", func(t *testing.T) {
		assert.Equal(t, 2, utils.GetFilledFieldsCount(&struct {
			name *string
			age  uint
		}{
			name: &name,
			age:  age,
		}))
	})

	t.Run("Получение количества заполненных полей структуры с не заполненную структурой не опциональных значений", func(t *testing.T) {
		assert.Equal(t, 0, utils.GetFilledFieldsCount(&struct {
			name string
			age  uint
		}{
			name: "",
			age:  0,
		}))
	})
}
