package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/muhafs/go-restful-api/app"
	"github.com/muhafs/go-restful-api/controlller"
	"github.com/muhafs/go-restful-api/exception"
	"github.com/muhafs/go-restful-api/helper"
	"github.com/muhafs/go-restful-api/repository"
	"github.com/muhafs/go-restful-api/service"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	ctgRepository := repository.NewCategoryRepository()
	ctgService := service.NewCategoryService(ctgRepository, db, validate)
	ctgController := controlller.NewCategoryController(ctgService)

	router := httprouter.New()

	router.GET("/api/categories", ctgController.FindAll)
	router.POST("/api/categories", ctgController.Create)
	router.GET("/api/categories/:categoryID", ctgController.FindOne)
	router.PUT("/api/categories/:categoryID", ctgController.Update)
	router.DELETE("/api/categories/:categoryID", ctgController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
