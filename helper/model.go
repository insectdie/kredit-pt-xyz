package helper

import (
	"insectdie/kredit-pt-xyz/model/domain"
	"insectdie/kredit-pt-xyz/model/web"
)

func ToKonsumenResponse(konsumen domain.Konsumen) web.KonsumenResponse {
	return web.KonsumenResponse{
		Nik:          konsumen.Nik,
		Full_name:    konsumen.Full_name,
		Legal_name:   konsumen.Legal_name,
		Tempat_lahir: konsumen.Tempat_lahir,
		Gaji:         konsumen.Gaji,
		Foto_ktp:     konsumen.Foto_ktp,
		Foto_selfie:  konsumen.Foto_selfie,
	}
}

func ToKonsumenResponses(konsumens []domain.Konsumen) []web.KonsumenResponse {
	var konsumenResponses []web.KonsumenResponse
	for _, konsumen := range konsumens {
		konsumenResponses = append(konsumenResponses, ToKonsumenResponse(konsumen))
	}
	return konsumenResponses
}
