package models

import (
	"database/sql"
	"mime/multipart"
	"time"
)

type Recipe struct {
	ID             uint          `json:"id,omitempty" db:"id"`
	Name           string        `json:"name" db:"name"`
	Description    string        `json:"description" db:"description"`
	PhotoURL       string        `json:"photoUrl,omitempty" db:"photo_url"`
	Ingredients    []string      `json:"ingredients" db:"ingredients"`
	Steps          []RecipeStep  `json:"steps" db:"steps"`
	TotalStepsTime time.Duration `json:"totalStepsTime,omitempty" db:"total_steps_time"`
	UpdatedAt      time.Time     `json:"updatedAt,omitempty" db:"updated_at"`
	DeletedAt      sql.NullTime  `json:"deletedAt,omitempty" db:"deleted_at"`
	CreatedAt      time.Time     `json:"createdAt,omitempty" db:"created_at"`
	Deleted        bool          `json:"deleted,omitempty" db:"deleted"`
}

type RecipeStep struct {
	Title       string        `json:"title" db:"title"`
	Description string        `json:"description" db:"description"`
	PhotoURL    string        `json:"photoUrl,omitempty" db:"photo_url"`
	Time        time.Duration `json:"time" db:"time"`
}

type SetRecipeStepPhoto struct {
	RecipeID    uint
	RecipeIndex uint
	Photo       *multipart.FileHeader
}

type SetRecipePhoto struct {
	RecipeID uint
	Photo    *multipart.FileHeader
}
