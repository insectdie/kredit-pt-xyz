package service

import (
	"context"
	"database/sql"
	"insectdie/kredit-pt-xyz/exception"
	"insectdie/kredit-pt-xyz/helper"
	"insectdie/kredit-pt-xyz/model/domain"
	"insectdie/kredit-pt-xyz/model/web"
	"insectdie/kredit-pt-xyz/repository"

	"github.com/go-playground/validator/v10"
)

type TransaksiServiceImpl struct {
	TransaksiRepository repository.TransaksiRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewTransaksiService(transaksiRepository repository.TransaksiRepository, DB *sql.DB, validate *validator.Validate) TransaksiService {
	return &TransaksiServiceImpl{
		TransaksiRepository: transaksiRepository,
		DB:                  DB,
		Validate:            validate,
	}
}

func (service *TransaksiServiceImpl) Create(ctx context.Context, request web.TransaksiCreateRequest) web.TransaksiResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transaksi := domain.Transaksi{
		Nik:         request.Nik,
		Otr:         request.Otr,
		Admin_fee:   request.Admin_fee,
		Jml_cicilan: request.Jml_cicilan,
		Jml_bunga:   request.Jml_bunga,
		Nama_asset:  request.Nama_asset,
	}

	transaksi = service.TransaksiRepository.Save(ctx, tx, transaksi)

	return helper.ToTransaksiResponse(transaksi)
}

func (service *TransaksiServiceImpl) Delete(ctx context.Context, transaksiNo_keluar int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transaksi, err := service.TransaksiRepository.FindById(ctx, tx, transaksiNo_keluar)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.TransaksiRepository.Delete(ctx, tx, transaksi)
}

func (service *TransaksiServiceImpl) FindById(ctx context.Context, transaksiNo_keluar int) web.TransaksiResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transaksi, err := service.TransaksiRepository.FindById(ctx, tx, transaksiNo_keluar)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToTransaksiResponse(transaksi)
}
