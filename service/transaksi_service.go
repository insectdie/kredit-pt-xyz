package service

import (
	"context"
	"insectdie/kredit-pt-xyz/model/web"
)

type TransaksiService interface {
	Create(ctx context.Context, request web.TransaksiCreateRequest) web.TransaksiResponse
	Delete(ctx context.Context, TransaksiId int)
	FindById(ctx context.Context, TransaksiId int) web.TransaksiResponse
}
