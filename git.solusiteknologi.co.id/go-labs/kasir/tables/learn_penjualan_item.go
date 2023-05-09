package tables

import (
	"git.solusiteknologi.co.id/goleaf/goleafcore/glentity"
	"github.com/shopspring/decimal"
)

type LearnPenjualanItem struct {
	PenjualanItemId int64           `json:penjualanItem`
	PenjualanId     int64           `json:penjualanId`
	ProductId       int64           `json:productId`
	Qty             decimal.Decimal `json:qty`
	Price           decimal.Decimal `json:price`

	glentity.BaseEntity
}
