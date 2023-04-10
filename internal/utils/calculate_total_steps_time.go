package utils

import (
	"time"

	"github.com/yeahyeahcore/recipes/internal/models"
)

func CalculateTotalStepsTime(steps []models.RecipeStep) time.Duration {
	totalTime := time.Duration(0)

	for _, step := range steps {
		totalTime += step.Time
	}

	return totalTime
}
