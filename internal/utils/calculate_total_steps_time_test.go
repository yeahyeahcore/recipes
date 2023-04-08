package utils_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"penahub.gitlab.yandexcloud.net/pena-services/accruals-service/internal/models"
	"penahub.gitlab.yandexcloud.net/pena-services/accruals-service/internal/utils"
)

func TestCalculateTotalStepsTime(t *testing.T) {
	t.Run("Calculate total steps time", func(t *testing.T) {
		assert.Equal(t, time.Minute*65, utils.CalculateTotalStepsTime([]models.RecipeStep{
			{Title: "step1", Time: time.Minute * 5},
			{Title: "step2", Time: time.Minute * 20},
			{Title: "step3", Time: time.Minute * 40},
		}))
	})

	t.Run("Calculate total steps with empty time", func(t *testing.T) {
		assert.Equal(t, time.Duration(0), utils.CalculateTotalStepsTime([]models.RecipeStep{
			{Title: "step1"},
			{Title: "step2"},
			{Title: "step3"},
		}))
	})
}
