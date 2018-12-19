package errors

// APIError is the model representing custom error
type APIError struct {
	Message       string `json:"message"`
	ResponseCode  int    `json:"-"`
	OriginalError string `json:"-"`
}

//FromError will add the original error in the APIError
func (f APIError) FromError(err error) APIError {
	f.OriginalError = err.Error()
	return f
}

func (f APIError) Error() string {
	return f.Message
}

// StatusCode return the proper status code
// Imported from kithttp
func (f APIError) StatusCode() int {
	return f.ResponseCode
}
