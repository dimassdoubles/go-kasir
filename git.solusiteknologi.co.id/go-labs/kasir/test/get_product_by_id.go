package test

import (
	"testing"

	"git.solusiteknologi.co.id/go-labs/kasir/dao/productdao"
	"git.solusiteknologi.co.id/goleaf/goleafcore"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glinit"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gltest"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

func TestGetProductById(t *testing.T) {
	gltest.TestDb(t, func(tx pgx.Tx) error {
		result, err := productdao.GetProductById(productdao.InputGetProductById{
			Tx: tx,
			ProductId: 1,
		})

		if err != nil {
			return err
		}

		logrus.Debug("SELECTED product : ", goleafcore.NewOrEmpty(result).PrettyString())

		return nil

	}, func(assert *gltest.Assert, tx pgx.Tx) interface{} {
		glinit.InitLog(glinit.LogConfig{
			LogFile:  "/log/get_product_by_id.log",
			LogLevel: glinit.DEFAULT_LOG_LEVEL,
		})

		return nil
	})
}
