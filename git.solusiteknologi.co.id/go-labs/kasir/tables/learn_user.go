package tables

import (
	"git.solusiteknologi.co.id/goleaf/goleafcore/glentity"
)

type LearnUser struct {
	UserId   int64  `json:userId`
	Username string `json:username`
	Password string `json:password`

	glentity.MasterEntity
}
