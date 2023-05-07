package service

import (
	"context"
	"insectdie/kredit-pt-xyz/model/web"
)

type KonsumenService interface {
	Create(ctx context.Context, request web.KonsumenCreateRequest) web.KonsumenResponse
	Update(ctx context.Context, request web.KonsumenUpdateRequest) web.KonsumenResponse
	Delete(ctx context.Context, konsumenId int)
	FindById(ctx context.Context, konsumenId int) web.KonsumenResponse
	FindAll(ctx context.Context) []web.KonsumenResponse
}
