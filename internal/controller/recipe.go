package controller

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"penahub.gitlab.yandexcloud.net/pena-services/accruals-service/internal/models"
	"penahub.gitlab.yandexcloud.net/pena-services/accruals-service/pkg/json"
)

type recipeService interface {
	Create(ctx context.Context, recipe *models.Recipe) (*models.Recipe, error)
	Update(ctx context.Context, recipe *models.Recipe) (*models.Recipe, error)
	GetByID(ctx context.Context, id uint) (*models.Recipe, error)
	GetAll(ctx context.Context) ([]models.Recipe, error)
	GetByIngredients(ctx context.Context, ingredients []string) ([]models.Recipe, error)
	GetByTotalStepsTime(ctx context.Context, totalTime time.Duration) ([]models.Recipe, error)
	SetRecipeStepPhoto(ctx context.Context, request *models.SetRecipeStepPhoto) error
	SetRecipePhoto(ctx context.Context, request *models.SetRecipePhoto) error
}

type RecipeController struct {
	logger  *logrus.Logger
	service recipeService
}

func NewRecipeController(logger *logrus.Logger, service recipeService) *RecipeController {
	return &RecipeController{
		logger:  logger,
		service: service,
	}
}

func (receiver *RecipeController) CreateRecipe(ctx echo.Context) error {
	request, err := json.Parse[models.Recipe](ctx.Request().Body)
	if err != nil {
		receiver.logger.Errorf("failed to parse request body on <CreateRecipe>: %w", err)
		return response(ctx, "failed to parse request body on <CreateRecipe>", err)
	}

	createdRecipe, err := receiver.service.Create(ctx.Request().Context(), request)
	if err != nil {
		return response(ctx, "failed to create recipe", err)
	}

	return ctx.JSON(http.StatusOK, createdRecipe)
}

func (receiver *RecipeController) UpdateRecipe(ctx echo.Context) error {
	request, err := json.Parse[models.Recipe](ctx.Request().Body)
	if err != nil {
		receiver.logger.Errorf("failed to parse request body on <UpdateRecipe>: %w", err)
		return response(ctx, "failed to parse request body on <UpdateRecipe>", err)
	}

	updatedRecipe, err := receiver.service.Update(ctx.Request().Context(), request)
	if err != nil {
		return response(ctx, "failed to update recipe", err)
	}

	return ctx.JSON(http.StatusOK, updatedRecipe)
}

func (receiver *RecipeController) GetAll(ctx echo.Context) error {
	recipes, err := receiver.service.GetAll(ctx.Request().Context())
	if err != nil {
		return response(ctx, "failed to get all recipes", err)
	}

	return ctx.JSON(http.StatusOK, recipes)
}

func (receiver *RecipeController) GetByID(ctx echo.Context) error {
	paramID := ctx.Param("id")

	id, err := strconv.ParseUint(paramID, 10, 64)
	if err != nil {
		return response(ctx, "recipe id must be uint", ErrInvalidParam)
	}

	recipes, err := receiver.service.GetByID(ctx.Request().Context(), uint(id))
	if err != nil {
		return response(ctx, "failed to get all recipes", err)
	}

	return ctx.JSON(http.StatusOK, recipes)
}

func (receiver *RecipeController) GetByIngredients(ctx echo.Context) error {
	ingredients := ctx.QueryParams()["ingredients"]

	recipes, err := receiver.service.GetByIngredients(ctx.Request().Context(), ingredients)
	if err != nil {
		return response(ctx, "failed to get recipes by ingredients", err)
	}

	return ctx.JSON(http.StatusOK, recipes)
}

func (receiver *RecipeController) GetByTotalStepsTime(ctx echo.Context) error {
	untilTotalTime, err := strconv.ParseInt(ctx.Param("totalTime"), 10, 64)
	if err != nil {
		return response(ctx, "recipe id must be uint", ErrInvalidParam)
	}

	recipes, err := receiver.service.GetByTotalStepsTime(ctx.Request().Context(), time.Duration(untilTotalTime))
	if err != nil {
		return response(ctx, "failed to get recipes by total recipes time", err)
	}

	return ctx.JSON(http.StatusOK, recipes)
}

func (receiver *RecipeController) SetRecipeStepPhoto(ctx echo.Context) error {
	recipeID, err := strconv.Atoi(ctx.Param("recipeID"))
	if err != nil {
		receiver.logger.Errorf("failed to parse recipe id: %w", err)
		return response(ctx, "invalid recipe id", ErrInvalidParam)
	}

	stepIndex, err := strconv.Atoi(ctx.Param("stepIndex"))
	if err != nil {
		receiver.logger.Errorf("failed to parse step index: %w", err)
		return response(ctx, "invalid step index", ErrInvalidParam)
	}

	file, err := ctx.FormFile("photo")
	if err != nil {
		receiver.logger.Errorf("failed to get file: %w", err)
		return response(ctx, "failed to get file", ErrInvalidParam)
	}

	if err := receiver.service.SetRecipeStepPhoto(ctx.Request().Context(), &models.SetRecipeStepPhoto{
		RecipeID:    uint(recipeID),
		RecipeIndex: uint(stepIndex),
		Photo:       file,
	}); err != nil {
		receiver.logger.Errorf("failed to set photo to recipe step: %w", err)
		return response(ctx, "failed to set photo to recipe step", err)
	}

	return response(ctx, "recipe step photo has been set", nil)
}

func (receiver *RecipeController) SetRecipePhoto(ctx echo.Context) error {
	recipeID, err := strconv.Atoi(ctx.Param("recipeID"))
	if err != nil {
		receiver.logger.Errorf("failed to parse recipe id: %w", err)
		return response(ctx, "invalid recipe id", ErrInvalidParam)
	}

	file, err := ctx.FormFile("photo")
	if err != nil {
		receiver.logger.Errorf("failed to get file: %w", err)
		return response(ctx, "failed to get file", ErrInvalidParam)
	}

	if err := receiver.service.SetRecipePhoto(ctx.Request().Context(), &models.SetRecipePhoto{
		RecipeID: uint(recipeID),
		Photo:    file,
	}); err != nil {
		receiver.logger.Errorf("failed to set photo to recipe: %w", err)
		return response(ctx, "failed to set photo to recipe", err)
	}

	return response(ctx, "recipe photo has been set", nil)
}
