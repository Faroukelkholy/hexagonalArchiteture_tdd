package output

import "hexagonalArchiteture_tdd/domain/model"

type IUserCrudsPort interface {
	CreateUser(user *model.User) error
	GetUser(id string) (*model.User, error)
}
