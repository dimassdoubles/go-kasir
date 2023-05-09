package test

import (
	"testing"

	"git.solusiteknologi.co.id/go-labs/kasir/dao/penjualanitemdao"
	"git.solusiteknologi.co.id/goleaf/goleafcore"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glinit"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gltest"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glutil"
	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

func TestAddPenjualanItem(t *testing.T) {
	gltest.TestDb(t, func(tx pgx.Tx) error {
		result, err := penjualanitemdao.Add(penjualanitemdao.InputAdd{
			Tx:            tx,
			PenjualanId:   1,
			ProductId:     2,
			Qty:           decimal.NewFromInt(13),
			Price:         decimal.NewFromInt(200),
			AuditUserId:   9,
			AuditDatetime: glutil.DateTimeNow(),
		})

		if err != nil {
			return err
		}

		assert := gltest.NewAssert(t)

		// productId tidak dapat dipastikan karena selalu beda menggunakan sequence
		assert.AssertEquals(1, result.PenjualanId)
		assert.AssertEquals(2, result.ProductId)
		assert.AssertEquals(decimal.NewFromInt(13), result.Qty)
		assert.AssertEquals(decimal.NewFromInt(200), result.Price)
		assert.AssertEquals(0, result.Version)

		logrus.Debug("ADDED penjualan item : ", goleafcore.NewOrEmpty(result).PrettyString())

		return nil

	}, func(assert *gltest.Assert, tx pgx.Tx) interface{} {
		glinit.InitLog(glinit.LogConfig{
			LogFile:  "/log/add_penjualan_item.log",
			LogLevel: glinit.DEFAULT_LOG_LEVEL,
		})

		return nil
	})
}
