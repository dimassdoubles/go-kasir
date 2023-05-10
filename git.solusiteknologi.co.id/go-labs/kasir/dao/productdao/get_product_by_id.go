package productdao

import (
	"errors"
	"fmt"

	"git.solusiteknologi.co.id/go-labs/kasir/tables"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldb"
	"github.com/jackc/pgx/v4"
)

type InputGetProductById struct {
	Tx        pgx.Tx
	ProductId int64
}

func GetProductById(input InputGetProductById) (*tables.LearnProduct, error) {
	result := tables.LearnProduct{}

	err := gldb.SelectOneQTx(input.Tx, *gldb.NewQBuilder().
		Add("SELECT * FROM ", tables.LEARN_PRODUCT, " ").
		Add("WHERE product_id=:productId").
		SetParam("productId", input.ProductId),
		&result,
	)

	if err != nil {
		return &result, errors.New(fmt.Sprint("error get product by id ", err))
	}

	return &result, nil
}
