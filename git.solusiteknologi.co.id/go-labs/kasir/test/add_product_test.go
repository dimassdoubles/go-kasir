package test

import (
	"testing"

	"git.solusiteknologi.co.id/go-labs/kasir/dao/productdao"
	"git.solusiteknologi.co.id/goleaf/goleafcore"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glinit"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gltest"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glutil"
	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

func TestAddProduct(t *testing.T) {
	gltest.TestDb(t, func(tx pgx.Tx) error {
		result, err := productdao.Add(productdao.InputAdd{
			Tx:            tx,
			ProductCode:   "P1111",
			ProductName:   "Indomie Rendang",
			Price:         decimal.NewFromInt(3000),
			AuditDatetime: glutil.DateTimeNow(),
			AuditUserId:   1,
		})

		if err != nil {
			return err
		}

		assert := gltest.NewAssert(t)

		// productId tidak dapat dipastikan karena selalu beda menggunakan sequence
		assert.AssertEquals("P1111", result.ProductCode)
		assert.AssertEquals("Indomie Rendang", result.ProductName)
		assert.AssertEquals(decimal.NewFromInt(3000), result.Price)
		assert.AssertEquals(0, result.Version)

		logrus.Debug("ADDED product : ", goleafcore.NewOrEmpty(result).PrettyString())

		return nil

	}, func(assert *gltest.Assert, tx pgx.Tx) interface{} {
		glinit.InitLog(glinit.LogConfig{
			LogFile:  "/log/add_product.log",
			LogLevel: glinit.DEFAULT_LOG_LEVEL,
		})

		return nil
	})
}
