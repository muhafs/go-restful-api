package repository

import (
	"context"
	"database/sql"

	"github.com/muhafs/go-restful-api/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, ctg domain.Category) domain.Category

	Update(ctx context.Context, tx *sql.Tx, ctg domain.Category) domain.Category

	Delete(ctx context.Context, tx *sql.Tx, ctg domain.Category)

	FindOne(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error)

	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
