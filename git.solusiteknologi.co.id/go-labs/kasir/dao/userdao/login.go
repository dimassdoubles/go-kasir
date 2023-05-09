package userdao

import (
	"errors"
	"fmt"

	"git.solusiteknologi.co.id/go-labs/kasir/tables"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldb"
	"github.com/jackc/pgx/v4"
)

type InputLogin struct {
	Tx       pgx.Tx
	Username string
	Password string
}

func Login(input InputLogin) (*tables.LearnUser, error) {
	result := tables.LearnUser{}

	err := gldb.SelectOneQTx(input.Tx, *gldb.NewQBuilder().
		Add("SELECT * FROM ", tables.LEARN_USER, " ").
		Add("WHERE username=:username ").
		Add("AND password=:password ").
		SetParam("username", input.Username).
		SetParam("password", input.Password),
		&result)

	if err != nil {
		return nil, errors.New(fmt.Sprint("error login ", err))
	}

	return &result, nil
}
