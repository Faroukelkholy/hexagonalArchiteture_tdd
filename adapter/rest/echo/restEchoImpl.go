package echo

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"hexagonalArchiteture_tdd/adapter/rest/echo/controllers"
	"hexagonalArchiteture_tdd/domain"
	"log"
	"sync"
)

var once sync.Once

type RestEchoImpl struct {
	router *echo.Echo
}

func (thisREcho *RestEchoImpl) InitAdapter() {
	log.Println("RestGinImpl")
	thisREcho.router = echo.New()
	//thisREcho.router.Use(middleware.CORS())
	////thisREcho.router.Use(middleware.CSRF())
	//thisREcho.router.Use(middleware.Logger())
	//thisREcho.router.Use(middleware.Gzip())
	//thisREcho.router.Use(middleware.Secure())
}

func (thisREcho *RestEchoImpl) Start(port int) error {
	pingController := controllers.NewPingController(thisREcho.router)
	pingController.Control()
	userController := controllers.NewUserController(thisREcho.router, domain.UserUseCase)
	userController.Control()
	return thisREcho.router.Start(fmt.Sprintf(":%d", port))
}
