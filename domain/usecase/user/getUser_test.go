package user

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"hexagonalArchiteture_tdd/domain/model"
	"hexagonalArchiteture_tdd/test/mocks/adapter/repositoryMock"
	"log"
	"testing"
)

//var (
//	createUserRepoMockfunc func(user *model.User) error
//	getUserRepoMockfunc    func(Id string) (*model.User, error)
//)
//
//
//type userCrudsRepoMock struct {
//}
//
//func (*userCrudsRepoMock) CreateUser(user *model.User) error {
//	return createUserRepoMockfunc(user)
//}
//
//func (*userCrudsRepoMock) GetUser(id string) (*model.User, error) {
//	return getUserRepoMockfunc(id)
//}

func TestCreateUserSuccess(t *testing.T) {
	userId := "d864c9cd-f0ef-46c2-8959-f6b7dc02aa46"

	createUserRepoMockfunc := func(user *model.User) error {
		return nil
	}
	userRepoMock := &repositoryMock.UserCrudsRepoMock{createUserRepoMockfunc, nil}
	getUserUseCase := NewCreateUserUseCase(userRepoMock)
	err := getUserUseCase.CreateUser(&model.User{
		Id:        userId,
		FirstName: "Farouk",
		LastName:  "Elkholy",
	})
	log.Println("err in  TestCreateUserSuccess :", err)
	assert.Nil(t, err)
}

func TestGetUserNotExist(t *testing.T) {
	userId := "d864c9cd-f0ef-46c2-8959-f6b7dc02aa46"

	getUserRepoMockfunc := func(Id string) (*model.User, error) {
		return nil, errors.New("user do not exist")
	}
	userRepoMock := &repositoryMock.UserCrudsRepoMock{nil, getUserRepoMockfunc}
	getUserUseCase := NewGetUserUseCase(userRepoMock)
	result, err := getUserUseCase.GetUser(userId)
	log.Println("result in TestGetUserNotExist :", result)
	log.Println("err in  TestGetUserNotExist :", err)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, "user do not exist", err.Error())
}

func TestGetUser(t *testing.T) {
	userId := "d864c9cd-f0ef-46c2-8959-f6b7dc02aa46"
	getUserRepoMockfunc := func(Id string) (*model.User, error) {
		return &model.User{
			Id:        userId,
			FirstName: "Farouk",
			LastName:  "Elkholy",
		}, nil
	}
	userRepoMock := &repositoryMock.UserCrudsRepoMock{nil, getUserRepoMockfunc}
	getUserUseCase := NewGetUserUseCase(userRepoMock)
	result, err := getUserUseCase.GetUser(userId)
	log.Println("result in  TestGetUser :", result)
	log.Println("err in TestGetUser :", err)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, userId, result.Id)
	assert.EqualValues(t, "Farouk", result.FirstName)
	assert.EqualValues(t, "Elkholy", result.LastName)
}
