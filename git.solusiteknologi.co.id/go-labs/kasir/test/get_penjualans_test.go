package test

import (
	"testing"

	"git.solusiteknologi.co.id/go-labs/kasir/dao/penjualandao"
	"git.solusiteknologi.co.id/goleaf/goleafcore"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glinit"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gltest"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

func TestGetPenjualans(t *testing.T) {
	gltest.TestDb(t, func(tx pgx.Tx) error {
		result, err := penjualandao.GetPenjualans(penjualandao.InputGet{Tx: tx})

		if err != nil {
			return err
		}

		logrus.Debug("SELECTED penjualans : ", goleafcore.NewOrEmpty(result).PrettyString())

		return nil

	}, func(assert *gltest.Assert, tx pgx.Tx) interface{} {
		glinit.InitLog(glinit.LogConfig{
			LogFile:  "/log/get_penjualans.log",
			LogLevel: glinit.DEFAULT_LOG_LEVEL,
		})

		return nil
	})
}
