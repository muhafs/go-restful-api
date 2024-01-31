package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/muhafs/go-restful-api/helper"
	"github.com/muhafs/go-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, ctg domain.Category) domain.Category {
	SQL := "INSERT INTO categories(name) VALUE(?)"

	result, err := tx.ExecContext(ctx, SQL, ctg.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	ctg.ID = int(id)
	return ctg
}
func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, ctg domain.Category) domain.Category {
	SQL := "UPDATE categories SET name = ? WHERE id = ?"

	_, err := tx.ExecContext(ctx, SQL, ctg.Name, ctg.ID)
	helper.PanicIfError(err)

	return ctg
}
func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, ctg domain.Category) {
	SQL := "DELETE FROM categories WHERE id = ?"

	_, err := tx.ExecContext(ctx, SQL, ctg.ID)
	helper.PanicIfError(err)
}
func (repository *CategoryRepositoryImpl) FindOne(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error) {
	SQL := "SELECT id, name from categories WHERE id = ?"

	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.ID, &category.Name)
		helper.PanicIfError(err)

		return category, nil
	} else {
		return category, errors.New("cateegory not found.")
	}
}
func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id, name FROM categories"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}

		err := rows.Scan(&category.ID, &category.Name)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}

	return categories
}
