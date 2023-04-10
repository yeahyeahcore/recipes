package utils

import (
	"sort"

	"github.com/yeahyeahcore/recipes/internal/models"
)

func SortRecipesByTotalStepsTime(recipes []models.Recipe) {
	if len(recipes) < 1 {
		return
	}

	sort.Slice(recipes, func(firstIndex, secondIndex int) bool {
		return recipes[firstIndex].TotalStepsTime < recipes[secondIndex].TotalStepsTime
	})
}
