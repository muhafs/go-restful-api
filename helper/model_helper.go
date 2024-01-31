package helper

import (
	"github.com/muhafs/go-restful-api/model/domain"
	"github.com/muhafs/go-restful-api/model/web"
)

func ToCategoryResponse(ctg domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		ID:   ctg.ID,
		Name: ctg.Name,
	}
}

func ToListCategoryResponse(ctgs []domain.Category) []web.CategoryResponse {
	var cateegoryResponses []web.CategoryResponse
	for _, ctg := range ctgs {

		cateegoryResponses = append(cateegoryResponses, ToCategoryResponse(ctg))
	}

	return cateegoryResponses
}
