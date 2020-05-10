package rest

import (
	"hexagonalArchiteture_tdd/adapter/rest/echo"
	"hexagonalArchiteture_tdd/port/output"
	"sync"
)

var (
	instance *serviceAdapter
	once     sync.Once
)

type serviceAdapter struct {
	Adapter output.IRestServicePort
}

func NewRestAdapter() *serviceAdapter {
	once.Do(func() {
		instance = &serviceAdapter{}
		instance.Adapter = new(echo.RestEchoImpl)
		instance.Adapter.InitAdapter()
	})
	return instance
}
