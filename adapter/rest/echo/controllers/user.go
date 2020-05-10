package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"hexagonalArchiteture_tdd/port/input/userPort"
	"net/http"
)

type UserController struct {
	router              *echo.Echo
	userDomainInputPort userPort.IUserDomainServicePort
}

func NewUserController(router *echo.Echo, userDomainInputPort userPort.IUserDomainServicePort) *UserController {
	return &UserController{
		router,
		userDomainInputPort,
	}
}

func (thisUC *UserController) Control() {
	thisUC.getUser()
}

func (thisUC *UserController) getUser() {
	thisUC.router.GET("/users/:user_id", func(c echo.Context) (err error) {
		fmt.Printf("user %v\n", c.Param("user_id"))
		userId := c.Param("user_id")
		fmt.Printf("userId %v\n", userId)
		result, err := thisUC.userDomainInputPort.GetUser(userId)
		if err != nil {
			fmt.Println("getuser UseCase", err)
			//restErr := errors.BadRequest("invalid json body")
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(200, result)
	})
}
