package service

import (
	"brilianpmw/go-rest-open-api/model/web"
	"context"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) web.ProductCreateResponse
	Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductCreateResponse
	Delete(ctx context.Context, productId int)
	FindById(ctx context.Context, productId int) (web.ProductCreateResponse, error)
	FindAll(ctx context.Context) []web.ProductCreateResponse
}
