package tables

import (
	"fmt"

	"git.solusiteknologi.co.id/goleaf/goleafcore/glentity"
	"github.com/shopspring/decimal"
)

type LearnProduct struct {
	ProductId   int64           `json:productId`
	ProductCode string          `json:productCode`
	ProductName string          `json:productName`
	Price       decimal.Decimal `json:price`

	glentity.MasterEntity
}

func (p LearnProduct) String() string {
	return fmt.Sprint("[", p.ProductCode, "] ", p.ProductName, "   ", p.Price)
}
