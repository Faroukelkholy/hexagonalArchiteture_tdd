package main

import (
	"hexagonalArchiteture_tdd/adapter/repository"
	"hexagonalArchiteture_tdd/adapter/rest"
	"hexagonalArchiteture_tdd/domain"
	"hexagonalArchiteture_tdd/utils"
	"log"
)

func main() {
	repositoryInstance := repository.NewDatabaseAdapter()
	domain.NewDomainServiceImpl(repositoryInstance.Adapter)
	restService := rest.NewRestAdapter()
	log.Println("portttt:", utils.GetEnvConfig().HTTP_PORT)
	err := restService.Adapter.Start(utils.GetEnvConfig().HTTP_PORT)
	if err != nil {
		log.Fatal(err)
	}
}
