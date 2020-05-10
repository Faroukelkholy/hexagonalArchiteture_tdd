package userPort

import (
	"hexagonalArchiteture_tdd/domain/model"
)

type IGetUser interface {
	GetUser(Id string) (*model.User, error)
}
