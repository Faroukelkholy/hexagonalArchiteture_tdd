package userPort

import "hexagonalArchiteture_tdd/domain/model"

type ICreateUser interface {
	CreateUser(user *model.User) error
}
