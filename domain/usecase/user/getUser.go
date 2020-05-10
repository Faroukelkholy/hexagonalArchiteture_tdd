package user

import (
	"fmt"
	"hexagonalArchiteture_tdd/domain/model"
	"hexagonalArchiteture_tdd/port/input/userPort"
	"hexagonalArchiteture_tdd/port/output"
)

type GetUserUseCase struct {
	repository     output.IDatabasePort
}

func NewGetUserUseCase (repository output.IDatabasePort) userPort.IGetUser {
	return &GetUserUseCase{
		repository,
	}

}

func (thisUser *GetUserUseCase) GetUser(id string) (*model.User, error) {
	fmt.Printf("Use Case proccessing GetUser %s\n", id)
	// query db
	//TODO: database impl
	result, err :=thisUser.repository.GetUser(id)
	//result, err := thisUser.repository.GetBilling(id, accountId)
	if err != nil {
		fmt.Printf("Use Case proccessing GetBillingById err %v\n", err)
		return nil, err
	}
	fmt.Printf("Use Case proccessing GetBillingById result %v\n", result)
	return result, nil
}

