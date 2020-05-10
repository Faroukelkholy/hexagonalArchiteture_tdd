package domain

import (
	"hexagonalArchiteture_tdd/domain/usecase/user"
	"hexagonalArchiteture_tdd/port/input/userPort"
	"hexagonalArchiteture_tdd/port/output"
)

var UserUseCase userPort.IUserDomainServicePort

type userUseCase struct {
	userPort.IGetUser
}

type DomainServiceImpl struct {
	databasePort output.IDatabasePort
}

func NewDomainServiceImpl(repository output.IDatabasePort) *DomainServiceImpl {
	dm := DomainServiceImpl{
		repository,
	}
	dm.initUseCases()
	return &dm
}

func (thisDS *DomainServiceImpl) initUseCases() {
	UserUseCase = &userUseCase{
		user.NewGetUserUseCase(thisDS.databasePort),
	}
}
