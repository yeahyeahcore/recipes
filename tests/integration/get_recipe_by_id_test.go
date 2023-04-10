//go:build integration

package integration_test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/yeahyeahcore/recipes/internal/models"
	"github.com/yeahyeahcore/recipes/pkg/client"
	"github.com/yeahyeahcore/recipes/pkg/env"
)

func TestGetRecipeByID(t *testing.T) {
	config, err := env.Parse[models.Config]("../../.env.test")

	assert.NoError(t, err)
	assert.NotNil(t, config)

	resty := resty.New()

	address := fmt.Sprintf("http://%s:%s/recipe", config.HTTP.Host, config.HTTP.Port)

	recipe, err := client.Get(fmt.Sprintf("%s/1", address), &client.RequestSettings[models.Recipe]{
		Driver: resty,
	})

	createdAt := time.Date(2023, time.April, 10, 4, 50, 5, 587280000, time.UTC)

	isNoErr := assert.Nil(t, err)
	isRecipeExist := assert.NotNil(t, recipe)

	if isNoErr && isRecipeExist {
		recipe.CreatedAt = createdAt
		recipe.UpdatedAt = createdAt

		assert.Equal(t, &models.Recipe{
			ID:          1,
			Name:        "Chocolate Chip Cookies",
			Description: "Classic chocolate chip cookies made with butter, sugar, eggs, flour, and chocolate chips.",
			PhotoURL:    sql.NullString{String: "https://example.com/chocolate-chip-cookies.jpg", Valid: true},
			Ingredients: []string{"butter", "sugar", "eggs", "flour", "chocolate chips"},
			Steps: []models.RecipeStep{
				{
					Title:       "Mix the batter",
					Description: "In a large bowl, cream together the butter and sugar until light and fluffy. Beat in the eggs, one at a time, then stir in the flour and chocolate chips.",
				},
				{
					Title:       "Bake the cookies",
					Description: "Drop tablespoonfuls of dough onto a baking sheet and bake at 375 degrees F (190 degrees C) for 10 to 12 minutes or until lightly browned.",
				},
			},
			TotalStepsTime: 1500000000000,
			UpdatedAt:      createdAt,
			DeletedAt:      sql.NullTime{Valid: false},
			CreatedAt:      createdAt,
			Deleted:        false,
		}, recipe)
	}
}
