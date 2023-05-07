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

type KonsumenServiceImpl struct {
	KonsumenRepository repository.KonsumenRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewKonsumenService(konsumenRepository repository.KonsumenRepository, DB *sql.DB, validate *validator.Validate) KonsumenService {
	return &KonsumenServiceImpl{
		KonsumenRepository: konsumenRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *KonsumenServiceImpl) Create(ctx context.Context, request web.KonsumenCreateRequest) web.KonsumenResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	konsumen := domain.Konsumen{
		Nik:           request.Nik,
		Full_name:     request.Full_name,
		Legal_name:    request.Legal_name,
		Tempat_lahir:  request.Tempat_lahir,
		Tanggal_lahir: request.Tanggal_lahir,
		Gaji:          request.Gaji,
		Foto_ktp:      request.Foto_ktp,
		Foto_selfie:   request.Foto_selfie,
	}

	konsumen = service.KonsumenRepository.Save(ctx, tx, konsumen)

	return helper.ToKonsumenResponse(konsumen)
}

func (service *KonsumenServiceImpl) Update(ctx context.Context, request web.KonsumenUpdateRequest) web.KonsumenResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	konsumen, err := service.KonsumenRepository.FindById(ctx, tx, request.Nik)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	konsumen.Gaji = request.Gaji
	konsumen.Foto_selfie = request.Foto_selfie

	konsumen = service.KonsumenRepository.Update(ctx, tx, konsumen)

	return helper.ToKonsumenResponse(konsumen)
}

func (service *KonsumenServiceImpl) Delete(ctx context.Context, konsumenNik int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	konsumen, err := service.KonsumenRepository.FindById(ctx, tx, konsumenNik)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.KonsumenRepository.Delete(ctx, tx, konsumen)
}

func (service *KonsumenServiceImpl) FindById(ctx context.Context, konsumenNik int) web.KonsumenResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	konsumen, err := service.KonsumenRepository.FindById(ctx, tx, konsumenNik)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToKonsumenResponse(konsumen)
}

func (service *KonsumenServiceImpl) FindAll(ctx context.Context) []web.KonsumenResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	konsumens := service.KonsumenRepository.FindAll(ctx, tx)

	return helper.ToKonsumenResponses(konsumens)
}
