package utils

import (
	"sort"

	"penahub.gitlab.yandexcloud.net/pena-services/accruals-service/internal/models"
)

func SortRecipesByTotalStepsTime(recipes []models.Recipe) {
	if len(recipes) < 1 {
		return
	}

	sort.Slice(recipes, func(firstIndex, secondIndex int) bool {
		return recipes[firstIndex].TotalStepsTime < recipes[secondIndex].TotalStepsTime
	})
}
