package user

import (
	"fmt"
	"hexagonalArchiteture_tdd/domain/model"
	"hexagonalArchiteture_tdd/port/input/userPort"
	"hexagonalArchiteture_tdd/port/output"
)

type GetUserUseCase struct {
	repository     output.IUserCrudsPort
}

func NewGetUserUseCase (repository output.IUserCrudsPort) userPort.IGetUser {
	return &GetUserUseCase{
		repository,
	}
}

func (thisUser *GetUserUseCase) GetUser(id string) (*model.User, error) {
	fmt.Printf("Use Case proccessing GetUser %s\n", id)
	result, err :=thisUser.repository.GetUser(id)
	if err != nil {
		fmt.Printf("Use Case proccessing GetUser err %v\n", err)
		return nil, err
	}
	fmt.Printf("Use Case proccessing user result %v\n", result)
	return result, nil
}

