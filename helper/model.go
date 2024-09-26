package helper

import (
	"iqbalatma/learn-go-pzn-restful-api/model/domain"
	"iqbalatma/learn-go-pzn-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
