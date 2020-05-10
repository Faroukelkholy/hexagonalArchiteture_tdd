package repository

import (
	"hexagonalArchiteture_tdd/adapter/repository/go_pg"
	"hexagonalArchiteture_tdd/port/output"
	"sync"
)

var (
	instance *databaseAdapter
	once     sync.Once
)

type databaseAdapter struct {
	Adapter output.IDatabasePort
}

func NewDatabaseAdapter() *databaseAdapter {
	once.Do(func() {
		instance = &databaseAdapter{}
		instance.Adapter = new(go_pg.DatabaseGoPgImpl)
		instance.Adapter.InitAdapter()
	})
	return instance
}
