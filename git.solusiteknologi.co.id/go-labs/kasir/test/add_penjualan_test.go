package test

import (
	"testing"

	"git.solusiteknologi.co.id/go-labs/kasir/dao/penjualandao"
	"git.solusiteknologi.co.id/goleaf/goleafcore"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glinit"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gltest"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glutil"
	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

func TestAddPenjualan(t *testing.T) {
	gltest.TestDb(t, func(tx pgx.Tx) error {
		result, err := penjualandao.Add(penjualandao.InputAdd{
			Tx:              tx,
			TotalPenjualan:  decimal.NewFromInt(120),
			TotalPembayaran: decimal.NewFromInt(200),
			TotalKembalian:  decimal.NewFromInt(80),
			AuditUserId:     -1,
			AuditDatetime:   glutil.DateTimeNow(),
		})

		if err != nil {
			return err
		}

		assert := gltest.NewAssert(t)

		// productId tidak dapat dipastikan karena selalu beda menggunakan sequence
		assert.AssertEquals(decimal.NewFromInt(120), result.TotalPenjualan)
		assert.AssertEquals(decimal.NewFromInt(200), result.TotalPembayaran)
		assert.AssertEquals(decimal.NewFromInt(80), result.TotalKembalian)
		assert.AssertEquals(0, result.Version)

		logrus.Debug("ADDED penjualan : ", goleafcore.NewOrEmpty(result).PrettyString())

		return nil

	}, func(assert *gltest.Assert, tx pgx.Tx) interface{} {
		glinit.InitLog(glinit.LogConfig{
			LogFile:  "/log/add_penjualan.log",
			LogLevel: glinit.DEFAULT_LOG_LEVEL,
		})

		return nil
	})
}
