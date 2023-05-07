package repository

import (
	"context"
	"database/sql"
	"insectdie/kredit-pt-xyz/model/domain"
)

type KonsumenRepository interface {
	Save(ctx context.Context, tx *sql.Tx, konsumen domain.Konsumen) domain.Konsumen
	Update(ctx context.Context, tx *sql.Tx, konsumen domain.Konsumen) domain.Konsumen
	Delete(ctx context.Context, tx *sql.Tx, konsumen domain.Konsumen)
	FindById(ctx context.Context, tx *sql.Tx, konsumenId int) (domain.Konsumen, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Konsumen
}
