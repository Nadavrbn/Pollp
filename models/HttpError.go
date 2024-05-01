package models

type HttpError struct {
	ErrorMessage   string
	HttpStatusCode int
}

func (error *HttpError) Error() string {
	return error.ErrorMessage
}
