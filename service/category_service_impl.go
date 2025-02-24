package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"iqbalatma/learn-go-pzn-restful-api/helper"
	"iqbalatma/learn-go-pzn-restful-api/model/domain"
	"iqbalatma/learn-go-pzn-restful-api/model/web"
	"iqbalatma/learn-go-pzn-restful-api/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryServiceImpl(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) *CategoryServiceImpl {
	return &CategoryServiceImpl{CategoryRepository: categoryRepository, DB: DB, Validate: validate}
}

func (service CategoryServiceImpl) Create(ctx context.Context, categoryCreateRequest web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(categoryCreateRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()

	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: categoryCreateRequest.Name,
	}

	service.CategoryRepository.Save(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()

	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	category.Name = request.Name

	service.CategoryRepository.Save(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()

	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)
	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()

	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()

	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	var categoriesResponses []web.CategoryResponse
	for _, category := range categories {
		categoriesResponses = append(categoriesResponses, helper.ToCategoryResponse(category))
	}

	return categoriesResponses
}
