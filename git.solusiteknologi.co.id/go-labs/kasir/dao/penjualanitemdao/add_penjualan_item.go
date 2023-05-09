package penjualanitemdao

import (
	"errors"
	"fmt"

	"git.solusiteknologi.co.id/go-labs/kasir/tables"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldb"
	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
)

type InputAdd struct {
	Tx            pgx.Tx
	PenjualanId   int64
	ProductId     int64
	Qty           decimal.Decimal
	Price         decimal.Decimal
	AuditUserId   int64
	AuditDatetime string
	Version       int64
}

func Add(input InputAdd) (*tables.LearnPenjualanItem, error) {
	result := tables.LearnPenjualanItem{}

	err := gldb.SelectOneQTx(input.Tx, *gldb.NewQBuilder().
		Add("INSERT INTO ", tables.LEARN_PENJUALAN_ITEM, " ( ").
		Add("penjualan_id, product_id, qty, price, ").
		Add("create_datetime, ").
		Add("update_datetime, create_user_id, update_user_id, version ").
		Add(") VALUES ( ").
		Add(":penjualanId, :productId, :qty, :price, ").
		Add(":datetime , ").
		Add(":datetime, :userId, :userId, :version ").
		Add(") RETURNING ").
		Add("penjualan_item_id, penjualan_id, product_id, qty, price, ").
		Add("create_datetime, ").
		Add("update_datetime, create_user_id, update_user_id, version ").
		SetParam("penjualanId", input.PenjualanId).
		SetParam("productId", input.ProductId).
		SetParam("qty", input.Qty).
		SetParam("price", input.Price).
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
