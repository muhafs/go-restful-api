package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/muhafs/go-restful-api/controlller"
	"github.com/muhafs/go-restful-api/exception"
)

func NewRouter(ctgController controlller.CategoryController) *httprouter.Router {
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

	return router
}
