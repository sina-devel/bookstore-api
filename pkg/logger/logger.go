package logger

type (
	Logger interface {
		Info(field LogField)
		Warning(field LogField)
		Error(field LogField)
	}

	LogField struct {
		Section  string
		Function string
		Params   interface{}
		Message  string
	}
)
