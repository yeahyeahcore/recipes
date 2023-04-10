package initialize

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/yeahyeahcore/recipes/internal/repository"
)

type RepositoriesDeps struct {
	Logger *logrus.Logger
	Pool   *pgxpool.Pool
}

type Repositories struct {
	RecipeRepository *repository.RecipeRepository
}

func NewRepositories(deps *RepositoriesDeps) *Repositories {
	return &Repositories{
		RecipeRepository: repository.NewRecipeRepository(deps.Pool, deps.Logger),
	}
}
