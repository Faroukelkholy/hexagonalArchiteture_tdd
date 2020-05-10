package go_pg

import (
	"fmt"
	"github.com/go-pg/pg/v9"
	"hexagonalArchiteture_tdd/port/output"
	"hexagonalArchiteture_tdd/utils"
)

//var lock sync.Mutex
var config = utils.GetEnvConfig()

type DatabaseGoPgImpl struct {
	pgConnection *pg.DB
	output.IUserCrudsPort
}

func (thisDBgoPg *DatabaseGoPgImpl) InitAdapter() {
	thisDBgoPg.pgConnection = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%v", config.DB_HOST, config.DB_PORT),
		User:     config.DB_USERNAME,
		Password: config.DB_PASSWORD,
		Database: config.DB_NAME,
	})
	// This logs the go_pg query on the database.
	//if config.DEBUG == "*" || config.DEBUG == "database" {
	//	thisDBgoPg.pgConnection.AddQueryHook(dbLogger{})
	//}
	thisDBgoPg.IUserCrudsPort = &UserCruds{Database: thisDBgoPg.pgConnection}

}

//type dbLogger struct{}
//
//func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
//	fmt.Println(q.FormattedQuery())
//	return c, nil
//}
//func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
//	fmt.Println(c)
//	fmt.Println(q.FormattedQuery())
//	return nil
//}
