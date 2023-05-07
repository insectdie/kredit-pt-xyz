package main

import (
	"insectdie/kredit-pt-xyz/app"
	"insectdie/kredit-pt-xyz/controller"
	"insectdie/kredit-pt-xyz/helper"
	"insectdie/kredit-pt-xyz/middleware"
	"insectdie/kredit-pt-xyz/repository"
	"insectdie/kredit-pt-xyz/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	konsumenRepository := repository.NewKonsumenRepository()
	konsumenService := service.NewKonsumenService(konsumenRepository, db, validate)
	konsumenController := controller.NewKonsumenController(konsumenService)

	transaksiRepository := repository.NewTransaksiRepository()
	transaksiService := service.NewTransaksiService(transaksiRepository, db, validate)
	transaksiController := controller.NewTransaksiController(transaksiService)

	router := app.NewRouter(konsumenController, transaksiController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
