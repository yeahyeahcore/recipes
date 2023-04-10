package service

import (
	"context"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/yeahyeahcore/recipes/internal/models"
	"github.com/yeahyeahcore/recipes/internal/utils"
	"github.com/yeahyeahcore/recipes/pkg/array"
)

type recipeRepository interface {
	Create(ctx context.Context, recipe *models.Recipe) (uint, error)
	Update(ctx context.Context, recipe *models.Recipe) error
	GetAll(ctx context.Context) ([]models.Recipe, error)
	GetByID(ctx context.Context, id uint) (*models.Recipe, error)
	GetByIngredients(ctx context.Context, ingredients []string) ([]models.Recipe, error)
	GetByTotalStepsTime(ctx context.Context, totalTime time.Duration) ([]models.Recipe, error)
	Delete(ctx context.Context, id uint) error
}

type RecipeServiceDeps struct {
	Logger           *logrus.Logger
	FileService      *FileService
	RecipeRepository recipeRepository
}

type RecipeService struct {
	logger           *logrus.Logger
	recipeRepository recipeRepository
	fileService      *FileService
}

func NewRecipeService(deps *RecipeServiceDeps) *RecipeService {
	return &RecipeService{
		logger:           deps.Logger,
		recipeRepository: deps.RecipeRepository,
		fileService:      deps.FileService,
	}
}

func (receiver *RecipeService) GetAll(ctx context.Context) ([]models.Recipe, error) {
	recipes, err := receiver.recipeRepository.GetAll(ctx)
	if err != nil {
		receiver.logger.Errorf("failed to get all recipes on <RecipeService>: %w", err)
		return []models.Recipe{}, err
	}

	utils.SortRecipesByTotalStepsTime(recipes)

	return recipes, nil
}

func (receiver *RecipeService) GetByID(ctx context.Context, id uint) (*models.Recipe, error) {
	return receiver.recipeRepository.GetByID(ctx, id)
}

func (receiver *RecipeService) GetByIngredients(ctx context.Context, ingredients []string) ([]models.Recipe, error) {
	recipes, err := receiver.recipeRepository.GetByIngredients(ctx, ingredients)
	if err != nil {
		receiver.logger.Errorf("failed to get recipes by ingredients on <RecipeService>: %w", err)
		return []models.Recipe{}, err
	}

	utils.SortRecipesByTotalStepsTime(recipes)

	return recipes, nil
}

func (receiver *RecipeService) GetByTotalStepsTime(ctx context.Context, totalTime time.Duration) ([]models.Recipe, error) {
	recipes, err := receiver.recipeRepository.GetByTotalStepsTime(ctx, totalTime)
	if err != nil {
		receiver.logger.Errorf("failed to get recipes by total steps time on <RecipeService>: %w", err)
		return []models.Recipe{}, err
	}

	utils.SortRecipesByTotalStepsTime(recipes)

	return recipes, nil
}

func (receiver *RecipeService) SetRecipePhoto(ctx context.Context, request *models.SetRecipePhoto) error {
	if request == nil {
		receiver.logger.Errorln("empty request on <SetRecipePhoto>")
		return ErrEmptyArgs
	}

	waitGroup := sync.WaitGroup{}

	recipe, err := receiver.GetByID(ctx, request.RecipeID)
	if err != nil {
		receiver.logger.Errorf("failed to get receiver by id on <SetRecipePhoto>: %w", err)
		return err
	}

	filepath, err := receiver.fileService.FormFilePath(request.Photo.Filename)
	if err != nil {
		receiver.logger.Errorf("failed to form file path on <SetRecipePhoto>: %w", err)
		return ErrInternal
	}

	recipe.PhotoURL = filepath

	waitGroup.Add(2)

	go func() {
		if err := receiver.recipeRepository.Update(ctx, recipe); err != nil {
			receiver.logger.Errorf("failed to update recipe on <SetRecipePhoto>: %w", err)
		}
		waitGroup.Done()
	}()

	go func() {
		if err := receiver.fileService.SaveFile(ctx, &models.SaveFile{
			Fileheader: request.Photo,
			Path:       filepath,
		}); err != nil {
			receiver.logger.Errorf("failed to save file on <SetRecipePhoto>: %w", err)
		}
	}()

	waitGroup.Wait()

	return nil
}

func (receiver *RecipeService) SetRecipeStepPhoto(ctx context.Context, request *models.SetRecipeStepPhoto) error {
	if request == nil {
		receiver.logger.Errorln("empty request on <SetRecipeStepPhoto>")
		return ErrEmptyArgs
	}

	waitGroup := sync.WaitGroup{}

	recipe, err := receiver.GetByID(ctx, request.RecipeID)
	if err != nil {
		receiver.logger.Errorf("failed to get receiver by id on <SetRecipeStepPhoto>: %w", err)
		return err
	}

	filepath, err := receiver.fileService.FormFilePath(request.Photo.Filename)
	if err != nil {
		receiver.logger.Errorf("failed to form file path on <SetRecipeStepPhoto>: %w", err)
		return ErrInternal
	}

	indexOfRecipeStep := array.IndexOf(recipe.Steps, func(step models.RecipeStep, index int, steps []models.RecipeStep) bool {
		return uint(index) == request.RecipeIndex
	})

	recipe.Steps[indexOfRecipeStep].PhotoURL = filepath

	waitGroup.Add(2)

	go func() {
		if err := receiver.recipeRepository.Update(ctx, recipe); err != nil {
			receiver.logger.Errorf("failed to update recipe on <SetRecipeStepPhoto>: %w", err)
		}
		waitGroup.Done()
	}()

	go func() {
		if err := receiver.fileService.SaveFile(ctx, &models.SaveFile{
			Fileheader: request.Photo,
			Path:       filepath,
		}); err != nil {
			receiver.logger.Errorf("failed to save file on <SetRecipeStepPhoto>: %w", err)
		}
	}()

	waitGroup.Wait()

	return nil
}

func (receiver *RecipeService) Create(ctx context.Context, recipe *models.Recipe) (*models.Recipe, error) {
	createdRecipeID, err := receiver.recipeRepository.Create(ctx, recipe)
	if err != nil {
		receiver.logger.Errorf("failed to create recipe on <Create> of <RecipeService>: %w", err)
		return nil, err
	}

	recipe.ID = createdRecipeID

	return recipe, nil
}

func (receiver *RecipeService) Update(ctx context.Context, recipe *models.Recipe) (*models.Recipe, error) {
	if err := receiver.recipeRepository.Update(ctx, recipe); err != nil {
		receiver.logger.Errorf("failed to update recipe on <Update> of <RecipeService>: %w", err)
		return nil, err
	}

	return recipe, nil
}
