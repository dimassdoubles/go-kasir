package productdao

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
	ProductCode   string          `json:productCode`
	ProductName   string          `json:productName`
	Price         decimal.Decimal `json:price`
	AuditDatetime string          `json:auditDatetime`
	AuditUserId   int64           `json:auditUserId`
}

func Add(input InputAdd) (*tables.LearnProduct, error) {
	result := tables.LearnProduct{}

	err := gldb.SelectOneQTx(input.Tx, *gldb.NewQBuilder().
		Add("INSERT INTO ", tables.LEARN_PRODUCT, " ( ").
		Add("product_code, product_name, price,").
		Add("create_datetime, ").
		Add("update_datetime, create_user_id, update_user_id, version, ").
		Add("active, active_datetime, non_active_datetime").
		Add(") VALUES ( ").
		Add(":productCode, :productName, :price,").
		Add(":datetime , ").
		Add(":datetime, :userId, :userId, :version, ").
		Add(":active, :datetime, :nonActiveDatetime").
		Add(") RETURNING ").
		Add("product_id, product_code, product_name, price,").
		Add("create_datetime, update_datetime, create_user_id, update_user_id,").
		Add("version, active, active_datetime, non_active_datetime").
		SetParam("productCode", input.ProductCode).
		SetParam("productName", input.ProductName).
		SetParam("price", input.Price).
		SetParam("datetime", input.AuditDatetime).
		SetParam("userId", input.AuditUserId).
		SetParam("active", "Y").
		SetParam("nonActiveDatetime", "").
		SetParam("version", 0),

		&result,
	)

	if err != nil {
		return &result, errors.New(fmt.Sprint("error add product ", err))
	}

	return &result, nil
}
