package echo

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"hexagonalArchiteture_tdd/adapter/rest/echo/controllers"
	"hexagonalArchiteture_tdd/domain"
	"log"
	"sync"
)

var (
	once sync.Once
	c    echo.Context
)

type RestEchoImpl struct {
	router *echo.Echo
}

func (thisREcho *RestEchoImpl) InitAdapter() {
	log.Println("RestGinImpl")
	thisREcho.router = echo.New()
}

func (thisREcho *RestEchoImpl) Start(port int) error {
	pingController := controllers.NewPingController(thisREcho.router)
	pingController.Control()
	userController := controllers.NewUserController(thisREcho.router, domain.UserUseCase)
	userController.Control(c)
	return thisREcho.router.Start(fmt.Sprintf(":%d", port))
}
