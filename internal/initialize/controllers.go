package initialize

import (
	"github.com/sirupsen/logrus"

	"github.com/yeahyeahcore/recipes/internal/controller"
)

type ControllersDeps struct {
	Logger   *logrus.Logger
	Services *Services
}

type Controllers struct {
	RecipeController *controller.RecipeController
}

func NewControllers(deps *ControllersDeps) *Controllers {
	return &Controllers{
		RecipeController: controller.NewRecipeController(deps.Logger, deps.Services.RecipeService),
	}
}
