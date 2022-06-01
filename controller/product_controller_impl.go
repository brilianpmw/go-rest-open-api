package controller

import (
	"brilianpmw/go-rest-open-api/helper"
	"brilianpmw/go-rest-open-api/model/web"
	"brilianpmw/go-rest-open-api/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}
func (controller *ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	productCreateRequest := web.ProductCreateRequest{}
	err := decoder.Decode(&productCreateRequest)
	helper.PanicIfError(err)
	productResponse := controller.ProductService.Create(request.Context(), productCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)

}
func (controller *ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	productUpdateRequest := web.ProductUpdateRequest{}
	err := decoder.Decode(&productUpdateRequest)
	helper.PanicIfError(err)
	id, err := strconv.Atoi(params.ByName("productId"))
	helper.PanicIfError(err)

	productUpdateRequest.Id = id
	productResponse := controller.ProductService.Update(request.Context(), productUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
func (controller *ProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	id, err := strconv.Atoi(params.ByName("productId"))
	helper.PanicIfError(err)

	controller.ProductService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
func (controller *ProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	id, err := strconv.Atoi(params.ByName("productId"))
	helper.PanicIfError(err)

	productResponse, err := controller.ProductService.FindById(request.Context(), id)
	// helper.PanicIfError(err)
	webResponse := web.WebResponse{}
	fmt.Println(err)
	if err != nil {
		webResponse = web.WebResponse{
			Code:    400,
			Status:  "error",
			Message: err.Error(),
		}
	} else {
		webResponse = web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   productResponse,
		}
	}
	// fmt.Println(productResponse)

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
func (controller *ProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	productResponses := controller.ProductService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponses,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
