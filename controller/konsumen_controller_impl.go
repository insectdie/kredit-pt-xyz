package controller

import (
	"insectdie/kredit-pt-xyz/helper"
	"insectdie/kredit-pt-xyz/model/web"
	"insectdie/kredit-pt-xyz/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type KonsumenControllerImpl struct {
	KonsumenService service.KonsumenService
}

func NewKonsumenController(konsumenService service.KonsumenService) KonsumenController {
	return &KonsumenControllerImpl{
		KonsumenService: konsumenService,
	}
}

func (controller *KonsumenControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	konsumenCreateRequest := web.KonsumenCreateRequest{}
	helper.ReadFromRequestBody(request, &konsumenCreateRequest)

	konsumenResponse := controller.KonsumenService.Create(request.Context(), konsumenCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   konsumenResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *KonsumenControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	konsumenUpdateRequest := web.KonsumenUpdateRequest{}
	helper.ReadFromRequestBody(request, &konsumenUpdateRequest)

	konsumenNik := params.ByName("konsumenNik")
	nik, err := strconv.Atoi(konsumenNik)
	helper.PanicIfError(err)

	konsumenUpdateRequest.Nik = nik

	konsumenResponse := controller.KonsumenService.Update(request.Context(), konsumenUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   konsumenResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *KonsumenControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	konsumenNik := params.ByName("konsumenNik")
	nik, err := strconv.Atoi(konsumenNik)
	helper.PanicIfError(err)

	controller.KonsumenService.Delete(request.Context(), nik)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *KonsumenControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	konsumenNik := params.ByName("konsumenNik")
	nik, err := strconv.Atoi(konsumenNik)
	helper.PanicIfError(err)

	konsumenResponse := controller.KonsumenService.FindById(request.Context(), nik)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   konsumenResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *KonsumenControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	konsumenResponses := controller.KonsumenService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   konsumenResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
