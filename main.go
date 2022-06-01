package main

import (
	"brilianpmw/go-rest-open-api/app"
	"brilianpmw/go-rest-open-api/controller"
	"brilianpmw/go-rest-open-api/helper"
	"brilianpmw/go-rest-open-api/repository"
	"brilianpmw/go-rest-open-api/service"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDb()
	productRepository := repository.NewProductRepository()

	productService := service.NewProductService(productRepository, db)
	productController := controller.NewProductController(productService)
	router := httprouter.New()

	router.GET("/api/product", productController.FindAll)
	router.GET("/api/product/:productId", productController.FindById)

	router.POST("/api/product", productController.Create)
	router.PUT("/api/product/:productId", productController.Update)
	router.DELETE("/api/product/:productId", productController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
	log.Print("running at port 3000")
}
