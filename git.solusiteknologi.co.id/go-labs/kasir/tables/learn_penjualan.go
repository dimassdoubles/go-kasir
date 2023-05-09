package tables

import (
	"git.solusiteknologi.co.id/goleaf/goleafcore/glentity"
	"github.com/shopspring/decimal"
)

type LearnPenjualan struct {
	PenjualanId     int64           `json:penjualanId`
	TotalPenjualan  decimal.Decimal `json:totalPenjualan`
	TotalPembayaran decimal.Decimal `json:totalPembayaran`
	TotalKembalian  decimal.Decimal `json:totalKembalian`

	glentity.BaseEntity
}
