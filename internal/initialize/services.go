package initialize

import (
	"github.com/sirupsen/logrus"

	"github.com/yeahyeahcore/recipes/internal/service"
)

type ServicesDeps struct {
	Logger              *logrus.Logger
	Repositories        *Repositories
	FileServiceDistPath string
}

type Services struct {
	FileService   *service.FileService
	RecipeService *service.RecipeService
}

func NewServices(deps *ServicesDeps) *Services {
	fileService := service.NewFileService(deps.Logger, deps.FileServiceDistPath)

	return &Services{
		RecipeService: service.NewRecipeService(&service.RecipeServiceDeps{
			Logger:           deps.Logger,
			FileService:      fileService,
			RecipeRepository: deps.Repositories.RecipeRepository,
		}),
	}
}
