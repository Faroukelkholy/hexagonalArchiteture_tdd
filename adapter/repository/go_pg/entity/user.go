package entity

type User struct {
	tableName    struct{}  `pg:"user"`
	Id           string    // `pg:"id,type:uuid,default:gen_random_uuid()"`
	FirstName    string    // `pg:"first_name"`
	LastName     string    // `pg:"last_name"`
}
