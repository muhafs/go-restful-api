package controlller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)

	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)

	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)

	FindOne(writer http.ResponseWriter, request *http.Request, params httprouter.Params)

	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
