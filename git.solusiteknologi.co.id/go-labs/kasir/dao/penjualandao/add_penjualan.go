package penjualandao

import (
	"errors"
	"fmt"

	"git.solusiteknologi.co.id/go-labs/kasir/tables"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldb"
	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
)

type InputAdd struct {
	Tx              pgx.Tx
	TotalPenjualan  decimal.Decimal
	TotalPembayaran decimal.Decimal
	TotalKembalian  decimal.Decimal
	AuditUserId     int64
	AuditDatetime   string
}

func Add(input InputAdd) (*tables.LearnPenjualan, error) {
	result := tables.LearnPenjualan{}

	err := gldb.SelectOneQTx(input.Tx, *gldb.NewQBuilder().
		Add("INSERT INTO ", tables.LEARN_PENJUALAN, " ( ").
		Add("total_penjualan, total_pembayaran, total_kembalian, ").
		Add("create_datetime, ").
		Add("update_datetime, create_user_id, update_user_id, version ").
		Add(") VALUES ( ").
		Add(":totalPenjualan, :totalPembayaran, :totalKembalian, ").
		Add(":datetime , ").
		Add(":datetime, :userId, :userId, :version ").
		Add(") RETURNING ").
		Add("penjualan_id, total_penjualan, total_pembayaran, total_kembalian, ").
		Add("create_datetime, ").
		Add("update_datetime, create_user_id, update_user_id, version ").
		SetParam("totalPenjualan", input.TotalPenjualan).
		SetParam("totalPembayaran", input.TotalPembayaran).
		SetParam("totalKembalian", input.TotalKembalian).
		SetParam("datetime", input.AuditDatetime).
		SetParam("userId", input.AuditUserId).
		SetParam("version", 0),
		&result,
	)

	if err != nil {
		return nil, errors.New(fmt.Sprint("error add penjualan ", err))
	}

	return &result, nil
}
