package utils

import (
	"time"

	"penahub.gitlab.yandexcloud.net/pena-services/accruals-service/internal/models"
)

func CalculateTotalStepsTime(steps []models.RecipeStep) time.Duration {
	totalTime := time.Duration(0)

	for _, step := range steps {
		totalTime += step.Time
	}

	return totalTime
}
