package constant

const (
	UnmarshalError string = "UnmarshalError"
	ParamError     string = "ParamError"
	ReadError      string = "ReadError"
	InternalError  string = "InternalError"
)

func GetErrorTypeMessage(errorType string) string {
	var message string

	switch errorType {
	case ReadError:
		message = "Unable to read body"
	case UnmarshalError:
		message = "Bad json"
	case ParamError:
		message = "Bad parameters"
	default:
		message = "Server error"
	}

	return message
}
