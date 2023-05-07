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

func ToTransaksiResponse(transaksi domain.Transaksi) web.TransaksiResponse {
	return web.TransaksiResponse{
		No_kontrak:  transaksi.No_kontrak,
		Nik:         transaksi.Nik,
		Otr:         transaksi.Otr,
		Admin_fee:   transaksi.Admin_fee,
		Jml_cicilan: transaksi.Jml_cicilan,
		Jml_bunga:   transaksi.Jml_bunga,
		Nama_asset:  transaksi.Nama_asset,
	}
}

func ToTransaksiResponses(transaksis []domain.Transaksi) []web.TransaksiResponse {
	var transaksiResponses []web.TransaksiResponse
	for _, transaksi := range transaksis {
		transaksiResponses = append(transaksiResponses, ToTransaksiResponse(transaksi))
	}
	return transaksiResponses
}
