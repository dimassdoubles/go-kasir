package penjualanitemdao

import (
	"errors"
	"fmt"

	"git.solusiteknologi.co.id/go-labs/kasir/tables"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldb"
	"github.com/jackc/pgx/v4"
)

type InputGetPjItemByPjId struct {
	Tx pgx.Tx
	PenjualanId int64
}

func GetPjItemByPjId(input InputGetPjItemByPjId) ([]*tables.LearnPenjualanItem, error) {
	results := []*tables.LearnPenjualanItem{}

	err := gldb.SelectQTx(input.Tx, *gldb.NewQBuilder().
		Add("SELECT * FROM ", tables.LEARN_PENJUALAN_ITEM, " ").
		Add("WHERE penjualan_id=:penjualanId").
		SetParam("penjualanId", input.PenjualanId),
		&results,
	)

	if err != nil {
		return nil, errors.New(fmt.Sprint("error get penjualan item by penjualan id ", err))
	}

	return results, nil
}