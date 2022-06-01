package service

import (
	"brilianpmw/go-rest-open-api/helper"
	"brilianpmw/go-rest-open-api/model/domain"
	"brilianpmw/go-rest-open-api/model/web"
	"brilianpmw/go-rest-open-api/repository"
	"context"
	"database/sql"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
}

func NewProductService(productRepository repository.ProductRepository, DB *sql.DB) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                DB,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) web.ProductCreateResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	product := domain.Product{
		Name: request.Name,
	}
	product = service.ProductRepository.Save(ctx, tx, product)
	return helper.ToProductResponse(product)
}
func (service *ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductCreateResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	product = domain.Product{
		Id:   request.Id,
		Name: request.Name,
	}
	product = service.ProductRepository.Update(ctx, tx, product)
	return helper.ToProductResponse(product)

}
func (service *ProductServiceImpl) Delete(ctx context.Context, productId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	helper.PanicIfError(err)

	helper.PanicIfError(err)
	service.ProductRepository.Delete(ctx, tx, product)

}
func (service *ProductServiceImpl) FindById(ctx context.Context, productId int) (web.ProductCreateResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	// helper.PanicIfError(err)
	if err != nil {
		return helper.ToProductResponse(product), err

	} else {
		return helper.ToProductResponse(product), nil

	}

}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []web.ProductCreateResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepository.FindAll(ctx, tx)
	var products_responses []web.ProductCreateResponse
	for _, product := range products {
		products_responses = append(products_responses, helper.ToProductResponse(product))
	}

	return products_responses
}
