package repository

import (
	"context"
	"database/sql"
	"errors"
	"insectdie/kredit-pt-xyz/helper"
	"insectdie/kredit-pt-xyz/model/domain"
)

type KonsumenRepositoryImpl struct {
}

func NewKonsumenRepository() KonsumenRepository {
	return &KonsumenRepositoryImpl{}
}

func (repository *KonsumenRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, konsumen domain.Konsumen) domain.Konsumen {
	SQL := "insert into konsumen( nik, full_name, legal_name, tempat_lahir, tanggal_lahir, gaji, foto_ktp, foto_selfie ) values (?, ?, ?, ?, STR_TO_DATE(?, '%d/%c/%Y'), ?, ?, ?)"
	tx.ExecContext(ctx, SQL,
		// "2222222222222222", "coba", "coba", "probolinggo", "23/05/1997", 10000000, "xxx", "xxx",
		konsumen.Nik,
		konsumen.Full_name,
		konsumen.Legal_name,
		konsumen.Tempat_lahir,
		konsumen.Tanggal_lahir,
		konsumen.Gaji,
		konsumen.Foto_ktp,
		konsumen.Foto_selfie,
	)

	return konsumen
}

func (repository *KonsumenRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, konsumen domain.Konsumen) domain.Konsumen {
	SQL := "update konsumen set gaji = ?, foto_selfie = ? where nik = ?"
	_, err := tx.ExecContext(ctx, SQL, konsumen.Gaji, konsumen.Foto_selfie, konsumen.Nik)
	helper.PanicIfError(err)

	return konsumen
}

func (repository *KonsumenRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, konsumen domain.Konsumen) {
	SQL := "delete from konsumen where nik = ?"
	_, err := tx.ExecContext(ctx, SQL, konsumen.Nik)
	helper.PanicIfError(err)
}

func (repository *KonsumenRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, konsumenNik int) (domain.Konsumen, error) {
	SQL := "select nik, full_name, legal_name, tempat_lahir, tanggal_lahir, gaji, foto_ktp, foto_selfie from konsumen where nik = ?"
	rows, err := tx.QueryContext(ctx, SQL, konsumenNik)
	helper.PanicIfError(err)
	defer rows.Close()

	konsumen := domain.Konsumen{}
	if rows.Next() {
		err := rows.Scan(&konsumen.Nik, &konsumen.Full_name, &konsumen.Legal_name, &konsumen.Tempat_lahir, &konsumen.Tanggal_lahir,
			&konsumen.Gaji, &konsumen.Foto_ktp, &konsumen.Foto_selfie)
		helper.PanicIfError(err)
		return konsumen, nil
	} else {
		return konsumen, errors.New("konsumen is not found")
	}
}

func (repository *KonsumenRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Konsumen {
	SQL := "select nik, full_name, legal_name, tempat_lahir, tanggal_lahir, gaji, foto_ktp, foto_selfie from konsumen order by nik"
	//, legal_name, tempat_lahir, tanggal_lahir, gaji, foto_ktp, foto_selfie
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var konsumens []domain.Konsumen
	for rows.Next() {
		konsumen := domain.Konsumen{}
		err := rows.Scan(&konsumen.Nik, &konsumen.Full_name, &konsumen.Legal_name, &konsumen.Tempat_lahir, &konsumen.Tanggal_lahir,
			&konsumen.Gaji, &konsumen.Foto_ktp, &konsumen.Foto_selfie)
		helper.PanicIfError(err)
		konsumens = append(konsumens, konsumen)
	}
	return konsumens
}
