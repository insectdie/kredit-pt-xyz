package repository

import (
	"context"
	"database/sql"
	"insectdie/kredit-pt-xyz/model/domain"
)

type TransaksiRepository interface {
	Save(ctx context.Context, tx *sql.Tx, Transaksi domain.Transaksi) domain.Transaksi
	Delete(ctx context.Context, tx *sql.Tx, Transaksi domain.Transaksi)
	FindById(ctx context.Context, tx *sql.Tx, TransaksiId int) (domain.Transaksi, error)
}
