package output

type IRestServicePort interface {
	InitAdapter()
	Start(port int) error
}

