package penjualandao

import (
	"errors"
	"fmt"

	"git.solusiteknologi.co.id/go-labs/kasir/tables"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldb"
	"github.com/jackc/pgx/v4"
)

type InputGet struct {
	Tx pgx.Tx
}

func GetPenjualans(input InputGet) ([]*tables.LearnPenjualan, error) {
	results := []*tables.LearnPenjualan{}

	err := gldb.SelectQTx(input.Tx, *gldb.NewQBuilder().
		Add("SELECT * FROM ", tables.LEARN_PENJUALAN),
		&results,
	)

	if err != nil {
		return nil, errors.New(fmt.Sprint("error get penjualan ", err))
	}

	return results, nil

}
