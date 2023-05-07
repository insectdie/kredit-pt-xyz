package controller

import (
	"insectdie/kredit-pt-xyz/helper"
	"insectdie/kredit-pt-xyz/model/web"
	"insectdie/kredit-pt-xyz/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type TransaksiControllerImpl struct {
	TransaksiService service.TransaksiService
}

func NewTransaksiController(transaksiService service.TransaksiService) TransaksiController {
	return &TransaksiControllerImpl{
		TransaksiService: transaksiService,
	}
}

func (controller *TransaksiControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transaksiCreateRequest := web.TransaksiCreateRequest{}
	helper.ReadFromRequestBody(request, &transaksiCreateRequest)

	transaksiResponse := controller.TransaksiService.Create(request.Context(), transaksiCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transaksiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransaksiControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transaksiNo_kontrak := params.ByName("transaksiNo_kontrak")
	nik, err := strconv.Atoi(transaksiNo_kontrak)
	helper.PanicIfError(err)

	controller.TransaksiService.Delete(request.Context(), nik)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransaksiControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transaksiNo_kontrak := params.ByName("transaksiNo_kontrak")
	nik, err := strconv.Atoi(transaksiNo_kontrak)
	helper.PanicIfError(err)

	transaksiResponse := controller.TransaksiService.FindById(request.Context(), nik)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transaksiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
