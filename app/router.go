package app

import (
	"insectdie/kredit-pt-xyz/controller"
	"insectdie/kredit-pt-xyz/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(konsumenController controller.KonsumenController, transaksiController controller.TransaksiController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/konsumens", konsumenController.FindAll)
	router.GET("/api/konsumens/:konsumenNik", konsumenController.FindById)
	router.POST("/api/konsumens", konsumenController.Create)
	router.PUT("/api/konsumens/:konsumenNik", konsumenController.Update)
	router.DELETE("/api/konsumens/:konsumenNik", konsumenController.Delete)

	router.GET("/api/transaksis/:transaksiNo_kontrak", transaksiController.FindById)
	router.POST("/api/transaksis", transaksiController.Create)
	router.DELETE("/api/transaksis/:transaksiNo_kontrak", transaksiController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
