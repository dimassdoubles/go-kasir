package productdao

import (
	"errors"
	"fmt"

	"git.solusiteknologi.co.id/go-labs/kasir/tables"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldb"
	"github.com/jackc/pgx/v4"
)

type InputGetProducts struct {
	Tx pgx.Tx
}

func GetProducts(input InputGetProducts) ([]*tables.LearnProduct, error) {
	results := make([]*tables.LearnProduct, 0)

	err := gldb.SelectQTx(input.Tx, *gldb.NewQBuilder().
		Add("SELECT * FROM ", tables.LEARN_PRODUCT, " ").
		Add("ORDER BY product_name ASC"),
		&results,
	)

	if err != nil {
		return nil, errors.New(fmt.Sprint("error get products ", err))
	}

	return results, nil
}
