package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"penahub.gitlab.yandexcloud.net/pena-services/accruals-service/internal/models"
	"penahub.gitlab.yandexcloud.net/pena-services/accruals-service/internal/utils"
)

type RecipeRepository struct {
	db     *pgxpool.Pool
	logger *logrus.Logger
}

func NewRecipeRepository(db *pgxpool.Pool, logger *logrus.Logger) *RecipeRepository {
	return &RecipeRepository{
		db:     db,
		logger: logger,
	}
}

func (receiver *RecipeRepository) Create(ctx context.Context, recipe *models.Recipe) (uint, error) {
	if recipe == nil {
		receiver.logger.Errorln("empty recipe on create")
		return 0, ErrEmptyArgs
	}

	insertedID := uint(0)

	if err := receiver.db.QueryRow(ctx, `
		INSERT INTO recipes(name, description, photo_url, ingredients, steps, total_steps_time)
		VALUES($1, $2, $3, $4, $5, $6)
		RETURNING id
		`,
		recipe.Name,
		recipe.Description,
		recipe.PhotoURL,
		pq.Array(recipe.Ingredients),
		pq.Array(recipe.Steps),
		utils.CalculateTotalStepsTime(recipe.Steps),
	).Scan(&insertedID); err != nil {
		receiver.logger.Errorf("failed to scan created recipe recipe: %w", err)
		return 0, ErrScanRecord
	}

	return insertedID, nil
}

func (receiver *RecipeRepository) Update(ctx context.Context, recipe *models.Recipe) error {
	if recipe == nil {
		receiver.logger.Errorln("empty recipe on update")
		return ErrEmptyArgs
	}

	_, err := receiver.db.Exec(ctx, `
		UPDATE recipes 
		SET name=$1, description=$2, photo_url=$3, ingredients=$4, steps=$5, total_steps_time=$6 updated_at=$7
		WHERE id=$8
	`,
		recipe.Name,
		recipe.Description,
		recipe.PhotoURL,
		pq.Array(recipe.Ingredients),
		pq.Array(recipe.Steps),
		utils.CalculateTotalStepsTime(recipe.Steps),
		recipe.UpdatedAt,
		recipe.ID,
	)
	if err != nil {
		receiver.logger.Errorf("failed to update recipe: %w", err)
		return ErrUpdateRecord
	}

	return nil
}

func (receiver *RecipeRepository) GetAll(ctx context.Context) ([]models.Recipe, error) {
	recipes := make([]models.Recipe, 0)

	rows, err := receiver.db.Query(ctx, "SELECT * FROM recipes WHERE NOT deleted")
	if err != nil {
		receiver.logger.Errorf("failed to query all recipes: %w", err)
		return recipes, ErrGetRecord
	}

	defer rows.Close()

	for rows.Next() {
		recipe := models.Recipe{}

		if err := rows.Scan(
			&recipe.ID,
			&recipe.Name,
			&recipe.Description,
			&recipe.PhotoURL,
			&recipe.Ingredients,
			&recipe.Steps,
			&recipe.TotalStepsTime,
			&recipe.UpdatedAt,
			&recipe.DeletedAt,
			&recipe.CreatedAt,
			&recipe.Deleted,
		); err != nil {
			receiver.logger.Errorf("failed to scan all recipes: %w", err)
			return recipes, ErrScanRecord
		}

		recipes = append(recipes, recipe)
	}

	if err := rows.Err(); err != nil {
		receiver.logger.Errorf("failed to iterate all recipes: %w", err)
		return recipes, ErrIterateRows
	}

	return recipes, nil
}

func (receiver *RecipeRepository) GetByID(ctx context.Context, id uint) (*models.Recipe, error) {
	recipe := models.Recipe{}

	if err := receiver.db.QueryRow(ctx, "SELECT * FROM recipe WHERE id = $1 AND deleted = false", id).Scan(
		&recipe.ID,
		&recipe.Name,
		&recipe.Description,
		&recipe.PhotoURL,
		&recipe.Ingredients,
		&recipe.Steps,
		&recipe.UpdatedAt,
		&recipe.DeletedAt,
		&recipe.CreatedAt,
		&recipe.Deleted,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			receiver.logger.Errorf("no recipe with <%d> id in db: %w", id, err)
			return nil, ErrNoRecord
		}

		receiver.logger.Errorf("failed to get recipe by <%d> id: %w", id, err)
		return nil, ErrGetRecord
	}

	return &recipe, nil
}

func (receiver *RecipeRepository) GetByIngredients(ctx context.Context, ingredients []string) ([]models.Recipe, error) {
	recipes := make([]models.Recipe, 0)

	if len(ingredients) < 1 {
		receiver.logger.Errorln("empty ingredients on <GetByIngredients>")
		return recipes, ErrEmptyArgs
	}

	rows, err := receiver.db.Query(ctx,
		"SELECT * FROM recipes WHERE NOT deleted AND ingredients @> $1",
		pq.Array(ingredients),
	)
	if err != nil {
		receiver.logger.Errorf("failed to get recipes by ingredients: %w", err)
		return recipes, ErrGetRecord
	}

	defer rows.Close()

	for rows.Next() {
		recipe := models.Recipe{}

		if err := rows.Scan(
			&recipe.ID,
			&recipe.Name,
			&recipe.Description,
			&recipe.PhotoURL,
			&recipe.Ingredients,
			&recipe.TotalStepsTime,
			&recipe.Steps,
			&recipe.UpdatedAt,
			&recipe.DeletedAt,
			&recipe.CreatedAt,
			&recipe.Deleted,
		); err != nil {
			receiver.logger.Errorf("failed to scan recipes by ingredients: %w", err)
			return recipes, ErrScanRecord
		}

		recipes = append(recipes, recipe)
	}

	if err := rows.Err(); err != nil {
		receiver.logger.Errorf("failed to iterate recipes by ingredients: %w", err)
		return recipes, ErrIterateRows
	}

	return recipes, nil
}

func (receiver *RecipeRepository) GetByTotalStepsTime(ctx context.Context, totalTime time.Duration) ([]models.Recipe, error) {
	recipes := make([]models.Recipe, 0)

	rows, err := receiver.db.Query(ctx, "SELECT * FROM recipes WHERE NOT deleted AND total_steps_time <= $1", totalTime)
	if err != nil {
		receiver.logger.Errorf("failed to get recipes by total steps time: %w", err)
		return recipes, ErrGetRecord
	}

	defer rows.Close()

	for rows.Next() {
		recipe := models.Recipe{}

		if err := rows.Scan(
			&recipe.ID,
			&recipe.Name,
			&recipe.Description,
			&recipe.PhotoURL,
			&recipe.Ingredients,
			&recipe.Steps,
			&recipe.UpdatedAt,
			&recipe.DeletedAt,
			&recipe.CreatedAt,
			&recipe.Deleted,
		); err != nil {
			receiver.logger.Errorf("failed to scan recipes by total steps time: %w", err)
			return recipes, ErrScanRecord
		}

		recipes = append(recipes, recipe)
	}

	if err := rows.Err(); err != nil {
		receiver.logger.Errorf("failed to iterate recipes by total steps time: %w", err)
		return recipes, ErrIterateRows
	}

	return recipes, nil
}

func (receiver *RecipeRepository) Delete(ctx context.Context, id uint) error {
	if _, err := receiver.db.Exec(ctx, `
		UPDATE recipes 
		SET deleted_at = now(), 
		deleted = true 
		WHERE id = $1 AND NOT deleted
	`, id); err != nil {
		receiver.logger.Errorf("failed to delete recipe by <%d> id: %w", id, err)
		return ErrDeleteRecord
	}

	return nil
}
