package utils_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"penahub.gitlab.yandexcloud.net/pena-services/accruals-service/internal/models"
	"penahub.gitlab.yandexcloud.net/pena-services/accruals-service/internal/utils"
)

func TestSortRecipesByTotalStepsTime(t *testing.T) {
	t.Run("Sort recipes by total steps time", func(t *testing.T) {
		recipes := []models.Recipe{
			{TotalStepsTime: time.Duration(60000)},
			{TotalStepsTime: time.Duration(40000)},
			{TotalStepsTime: time.Duration(80000)},
		}

		utils.SortRecipesByTotalStepsTime(recipes)

		assert.Equal(t,
			[]models.Recipe{
				{TotalStepsTime: time.Duration(40000)},
				{TotalStepsTime: time.Duration(60000)},
				{TotalStepsTime: time.Duration(80000)},
			},
			recipes,
		)
	})

	t.Run("Sort empty recipes by total steps time", func(t *testing.T) {
		recipes := []models.Recipe{}

		utils.SortRecipesByTotalStepsTime(recipes)

		assert.Equal(t, []models.Recipe{}, recipes)
	})
}
