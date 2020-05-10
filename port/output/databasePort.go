package output

type IDatabasePort interface {
	InitAdapter()
	IUserCrudsPort
}
