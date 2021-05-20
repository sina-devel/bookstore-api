package contract

type (
	HttpServer interface {
		Start(port int) error
	}
)
