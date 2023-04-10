//go:build integration

package integration_test

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"

	"github.com/yeahyeahcore/recipes/internal/models"
	"github.com/yeahyeahcore/recipes/pkg/client"
	"github.com/yeahyeahcore/recipes/pkg/env"
)

func TestGetAllRecipes(t *testing.T) {
	config, err := env.Parse[models.Config]("../../.env.test")

	assert.NoError(t, err)
	assert.NotNil(t, config)

	resty := resty.New()

	address := fmt.Sprintf("http://%s:%s/recipe", config.HTTP.Host, config.HTTP.Port)

	recipes, err := client.Get(fmt.Sprintf("%s/all", address), &client.RequestSettings[[]models.Recipe]{
		Driver: resty,
	})

	assert.Nil(t, err)
	assert.NotNil(t, recipes)
}
