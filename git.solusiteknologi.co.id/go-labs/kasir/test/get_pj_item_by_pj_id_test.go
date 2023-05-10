package test

import (
	"testing"

	"git.solusiteknologi.co.id/go-labs/kasir/dao/penjualanitemdao"
	"git.solusiteknologi.co.id/goleaf/goleafcore"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glinit"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gltest"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

func TestGetPjItemByPjId(t *testing.T) {
	gltest.TestDb(t, func(tx pgx.Tx) error {
		result, err := penjualanitemdao.GetPjItemByPjId(penjualanitemdao.InputGetPjItemByPjId{
			Tx: tx,
			PenjualanId: 2,
		})

		if err != nil {
			return err
		}

		logrus.Debug("SELECTED penjualan items : ", goleafcore.NewOrEmpty(result).PrettyString())

		return nil

	}, func(assert *gltest.Assert, tx pgx.Tx) interface{} {
		glinit.InitLog(glinit.LogConfig{
			LogFile:  "/log/get_pj_item_by_pj_id.log",
			LogLevel: glinit.DEFAULT_LOG_LEVEL,
		})

		return nil
	})
}
