package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/muhafs/go-restful-api/app"
	"github.com/muhafs/go-restful-api/controlller"
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

	// setup router
	router := app.NewRouter(ctgController)

	// setup server connection address
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	// start the application
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
