package repositoryMock

import (
	"hexagonalArchiteture_tdd/domain/model"
)


type UserCrudsRepoMock struct {
	CreateUserRepoMockfunc func(user *model.User) error
	GetUserRepoMockfunc    func(Id string) (*model.User, error)
}

func (thisUc *UserCrudsRepoMock) CreateUser(user *model.User) error {
	return thisUc.CreateUserRepoMockfunc(user)
}

func (thisUc *UserCrudsRepoMock) GetUser(id string) (*model.User, error) {
	return thisUc.GetUserRepoMockfunc(id)
}

