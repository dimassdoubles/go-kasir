package test

import (
	"testing"

	"git.solusiteknologi.co.id/go-labs/kasir/dao/userdao"
	"git.solusiteknologi.co.id/goleaf/goleafcore"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glinit"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gltest"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

func TestLogin(t *testing.T) {
	gltest.TestDb(t, func(tx pgx.Tx) error {
		user, err := userdao.Login(userdao.InputLogin{
			Tx:       tx,
			Username: "sts",
			Password: "sts123",
		})

		if err != nil {
			return err
		}

		logrus.Debug("LOGGEDIN user : ", goleafcore.NewOrEmpty(user).PrettyString())

		assert := gltest.NewAssert(t)

		assert.AssertEquals("sts", user.Username)
		assert.AssertEquals("sts123", user.Password)

		return nil

	}, func(assert *gltest.Assert, tx pgx.Tx) interface{} {
		glinit.InitLog(glinit.LogConfig{
			LogFile:  "/log/kasir.log",
			LogLevel: glinit.DEFAULT_LOG_LEVEL,
		})

		return nil
	})
}
