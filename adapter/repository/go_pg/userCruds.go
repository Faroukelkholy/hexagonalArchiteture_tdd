package go_pg

import (
	"errors"
	"fmt"
	"github.com/go-pg/pg/v9"
	"hexagonalArchiteture_tdd/adapter/repository/go_pg/entity"
	"hexagonalArchiteture_tdd/domain/model"
)

type UserCruds struct {
	Database *pg.DB
}

func (thisCS *UserCruds) CreateUser(user *model.User) error {
	entityModel := &entity.User{
		Id:           user.Id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
	}
	res, err := thisCS.Database.Model(entityModel).OnConflict("DO NOTHING").Insert()
	if err != nil {
		fmt.Println("error creating billing", err)
		return err
	}
	if res.RowsAffected() > 0 {
		fmt.Println("created")
		return nil

	} else {
		fmt.Println("did nothing")
		return errors.New("There a user  already")
	}
}



func (thisCS *UserCruds) GetUser(id string) (*model.User, error) {
	enitiyQueried := &entity.User{}

	err := thisCS.Database.
		Model(enitiyQueried).
		Where("id=? ", id).
		First()
	if err != nil {
		fmt.Printf("error UpdateUser billing %v\n", err)
		if err.Error() == "pg: no rows in result set" {
			return nil, err
		}
		return nil, err
	}
	fmt.Println(enitiyQueried)

	entityModel := &model.User{
		Id:           enitiyQueried.Id,
		FirstName:    enitiyQueried.FirstName,
		LastName:     enitiyQueried.LastName,
	}

	return entityModel, nil
}

