package user

import (
	"fmt"
	"hexagonalArchiteture_tdd/domain/model"
	"hexagonalArchiteture_tdd/port/input/userPort"
	"hexagonalArchiteture_tdd/port/output"
)

type CreateUserUseCase struct {
	repository output.IUserCrudsPort
}

func NewCreateUserUseCase(repository output.IUserCrudsPort) userPort.ICreateUser {
	return &CreateUserUseCase{
		repository,
	}
}

func (thisUser *CreateUserUseCase) CreateUser(user *model.User) error {
	fmt.Printf("Use Case proccessing CreateUser %s\n", user)
	err := thisUser.repository.CreateUser(user)
	//result, err := thisUser.repositoryMock.CreateBilling(id, accountId)
	if err != nil {
		fmt.Printf("Use Case proccessing CreateUser err %v\n", err)
		return err
	}
	return nil
}
