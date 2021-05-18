package contract

type (
	HttpServer interface {
		Start(port uint) error
	}
)
