package helper

import (
	"brilianpmw/go-rest-open-api/model/domain"
	"brilianpmw/go-rest-open-api/model/web"
)

func ToProductResponse(product domain.Product) web.ProductCreateResponse {
	data := web.ProductCreateResponse{
		Id:   product.Id,
		Name: product.Name,
	}
	return data
}
