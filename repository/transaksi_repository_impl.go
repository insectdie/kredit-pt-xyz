package repository

import (
	"context"
	"database/sql"
	"errors"
	"insectdie/kredit-pt-xyz/helper"
	"insectdie/kredit-pt-xyz/model/domain"
	"log"
	"strconv"
)

type TransaksiRepositoryImpl struct {
}

func NewTransaksiRepository() TransaksiRepository {
	return &TransaksiRepositoryImpl{}
}

func (repository *TransaksiRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, transaksi domain.Transaksi) domain.Transaksi {
	var limit int
	limitSQL := "SELECT nvl(limit_pinjaman,0) - (SELECT nvl(SUM(otr), 0) FROM `transaksi` WHERE nik = a.nik) as limit_pinjaman FROM `konsumen_limit` a WHERE nik = ? and bulan = (SELECT TIMESTAMPDIFF(MONTH,  created_datetime, DATE_SUB(SYSDATE(), INTERVAL -1 MONTH)) FROM `konsumen` WHERE nik = a.nik)"
	err := tx.QueryRowContext(ctx, limitSQL, transaksi.Nik).Scan(&limit)
	helper.PanicIfError(err)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with NIK %d\n", transaksi.Nik)
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		if transaksi.Otr > limit {
			log.Printf("Maximum limit is %q \n", strconv.Itoa(limit))
			return transaksi
		}
	}

	SQL := "insert into transaksi( nik, otr, admin_fee,	jml_cicilan, jml_bunga,	nama_asset ) values (?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL,
		transaksi.Nik,
		transaksi.Otr,
		transaksi.Admin_fee,
		transaksi.Jml_cicilan,
		transaksi.Jml_bunga,
		transaksi.Nama_asset,
	)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	transaksi.No_kontrak = strconv.FormatInt(id, 10)

	return transaksi
}

func (repository *TransaksiRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, transaksi domain.Transaksi) {
	SQL := "delete from transaksi where no_kontrak = ?"
	_, err := tx.ExecContext(ctx, SQL, transaksi.No_kontrak)
	helper.PanicIfError(err)
}

func (repository *TransaksiRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, transaksiNo_kontrak int) (domain.Transaksi, error) {
	SQL := "select no_kontrak, nik, otr, admin_fee,	jml_cicilan, jml_bunga,	nama_asset from transaksi where no_kontrak = ?"
	rows, err := tx.QueryContext(ctx, SQL, transaksiNo_kontrak)
	helper.PanicIfError(err)
	defer rows.Close()

	transaksi := domain.Transaksi{}
	if rows.Next() {
		err := rows.Scan(&transaksi.No_kontrak, &transaksi.Nik, &transaksi.Otr, &transaksi.Admin_fee, &transaksi.Jml_cicilan,
			&transaksi.Jml_bunga, &transaksi.Nama_asset)
		helper.PanicIfError(err)
		return transaksi, nil
	} else {
		return transaksi, errors.New("transaksi is not found")
	}
}
