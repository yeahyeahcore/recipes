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

	resty.BaseURL = fmt.Sprintf("%s:%s", config.HTTP.Host, config.HTTP.Port)

	recipes, err := client.Get("/all", &client.RequestSettings[[]models.Recipe]{
		Driver: resty,
	})

	assert.Nil(t, err)
	assert.NotNil(t, recipes)
}
