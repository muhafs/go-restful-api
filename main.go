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
	"github.com/muhafs/go-restful-api/middleware"
	"github.com/muhafs/go-restful-api/repository"
	"github.com/muhafs/go-restful-api/service"
)

func main() {
	// make database coonecttion
	db := app.NewDB()
	// make validator object
	validate := validator.New()

	// make repository
	ctgRepository := repository.NewCategoryRepository()
	// make service
	ctgService := service.NewCategoryService(ctgRepository, db, validate)
	// make controller
	ctgController := controlller.NewCategoryController(ctgService)

	// make router object
	router := httprouter.New()

	// make route endpoints
	router.GET("/api/categories", ctgController.FindAll)
	router.POST("/api/categories", ctgController.Create)
	router.GET("/api/categories/:categoryID", ctgController.FindOne)
	router.PUT("/api/categories/:categoryID", ctgController.Update)
	router.DELETE("/api/categories/:categoryID", ctgController.Delete)

	// set global error handler
	router.PanicHandler = exception.ErrorHandler

	// setup server connection address
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	// start the application
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
