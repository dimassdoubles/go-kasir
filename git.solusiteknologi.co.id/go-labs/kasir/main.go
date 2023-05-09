package main

import (
	"fmt"

	"git.solusiteknologi.co.id/go-labs/kasir/cashier"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldb"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glinit"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	// init log
	glinit.InitLog(glinit.LogConfig{
		LogFile:  "log/gopostgres.log",
		LogLevel: glinit.DEFAULT_LOG_LEVEL,
	})

	glinit.InitDb(glinit.DbConfig{
		User:              "sts",
		Password:          "Awesome123!",
		Port:              5432,
		Host:              "localhost",
		Name:              "go_kasir",
		ApplicationName:   "GoKasir",
		PoolMaxConnection: 1,
		PoolMinConnection: 1,
	})

	for {
		err := gldb.BeginTrx(func(trx pgx.Tx) error {
			cashier := &cashier.Cashier{
				Tx: trx,
			}

			err := cashier.Start()
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println()
			fmt.Println()

			return nil

		})

		if err != nil {
			logrus.Debug("Error connection occured : ", err)
		}
	}

}
