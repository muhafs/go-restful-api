package service

import (
	"context"

	"github.com/muhafs/go-restful-api/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, req web.CategoryCreateRequest) web.CategoryResponse

	Update(ctx context.Context, req web.CategoryUpdateRequest) web.CategoryResponse

	Delete(ctx context.Context, categoryID int)

	FindOne(ctx context.Context, categoryID int) web.CategoryResponse

	FindAll(ctx context.Context) []web.CategoryResponse
}
