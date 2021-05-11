package log

type (
	Logger interface {
		Info(field *Field)
		Warning(field *Field)
		Error(field *Field)
	}

	Field struct {
		Section  string
		Function string
		Params   map[string]interface{}
		Message  string
	}
)
