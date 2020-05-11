package controllers

import (
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"hexagonalArchiteture_tdd/domain"
	"hexagonalArchiteture_tdd/domain/model"
	"hexagonalArchiteture_tdd/domain/usecase/user"
	"hexagonalArchiteture_tdd/port/input/userPort"
	"hexagonalArchiteture_tdd/test/mocks/adapter/repositoryMock"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	userId = "d864c9cd-f0ef-46c2-8959-f6b7dc02aa46"
)

var (
	router                 *echo.Echo
)

type userUseCase struct {
	userPort.IGetUser
	userPort.ICreateUser
}


func TestGetUserControllerUserNotExist(t *testing.T) {
	getUserRepoMockfunc := func(Id string) (*model.User, error) {
		return nil, errors.New("user do not exist")
	}
	userRepoMock := &repositoryMock.UserCrudsRepoMock{nil, getUserRepoMockfunc}
	getUserUseCase := user.NewGetUserUseCase(userRepoMock)
	result, err := getUserUseCase.GetUser(userId)
	log.Println("result in TestGetUserControllerUserNotExist :", result)
	log.Println("err in  TestGetUserControllerUserNotExist :", err)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, "user do not exist", err.Error())
}

func TestGetUserController(t *testing.T) {
	//init repo mock
	getUserRepoMockfunc := func(Id string) (*model.User, error) {
		return &model.User{
			Id:        userId,
			FirstName: "Farouk",
			LastName:  "Elkholy",
		}, nil
	}
	userRepoMock := &repositoryMock.UserCrudsRepoMock{nil, getUserRepoMockfunc}

	//init use Case
	getUserUseCase := user.NewGetUserUseCase(userRepoMock)
	userUseCase := &userUseCase{getUserUseCase, nil}
	domain.UserUseCase = userUseCase

	// Setup context mock for controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:user_id")
	c.SetParamNames("user_id")
	c.SetParamValues(userId)

	//execute controller
	userControl := NewUserController(router, domain.UserUseCase)

	// Assertions
	if assert.NoError(t, userControl.getUser(c)) {
		var user model.User
		err := json.Unmarshal(rec.Body.Bytes(), &user)
		log.Println("user in  TestGetUserController :", user)
		log.Println("err in TestGetUserController :", err)

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.EqualValues(t, userId, user.Id)
		assert.EqualValues(t, "Farouk", user.FirstName)
		assert.EqualValues(t, "Elkholy", user.LastName)
	}

}
