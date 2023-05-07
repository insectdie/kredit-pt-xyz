package app

import (
	"insectdie/kredit-pt-xyz/controller"
	"insectdie/kredit-pt-xyz/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(konsumenController controller.KonsumenController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/konsumens", konsumenController.FindAll)
	router.GET("/api/konsumens/:konsumenNik", konsumenController.FindById)
	router.POST("/api/konsumens", konsumenController.Create)
	router.PUT("/api/konsumens/:konsumenNik", konsumenController.Update)
	router.DELETE("/api/konsumens/:konsumenNik", konsumenController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
