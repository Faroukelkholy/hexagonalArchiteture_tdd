package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"hexagonalArchiteture_tdd/domain/model"
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

func (thisUC *UserController) Control(c echo.Context) {
	thisUC.router.GET("/users/:user_id", thisUC.getUser)
	thisUC.router.POST("/users", thisUC.createUser)
}

func (thisUC *UserController) getUser(c echo.Context) error {
	fmt.Printf("userid in controller %v\n", c.Param("user_id"))
	userId := c.Param("user_id")
	result, err := thisUC.userDomainInputPort.GetUser(userId)
	if err != nil {
		fmt.Println("getuser UseCase", err)
		//restErr := errors.BadRequest("invalid json body")
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(200, result)
}

func (thisUC *UserController) createUser(c echo.Context) error {
	fmt.Printf("createUser controller \n")
	var user model.User
	if err := c.Bind(&user); err != nil {
		fmt.Println("error unmarshiling :",err)
		return c.JSON(http.StatusBadRequest, err)
	}

	err := thisUC.userDomainInputPort.CreateUser(&user)
	if err != nil {
		fmt.Println("createUser UseCase", err)
		if err.Error() == "user already exists" {
			return c.JSON(http.StatusConflict, err)
		}
		//restErr := errors.BadRequest("invalid json body")
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(201, map[string]interface{}{
		"message": "Success",
	})
}
