package utils

// Centralized Error Struct Used in CustomHTTPHandlerMiddlewareEcho in building-block/v2
type (
	// RequestError wrapper for error response
	RequestError struct {
		Error interface{} `json:"error"`
	}

	// ErrorStruct for validator
	ErrorStructValidator struct {
		Field  string `json:"field"`
		Reason string `json:"reason"`
	}

	// RequestErrors for multiple validator errors
	RequestErrors struct {
		StatusCode int                    `json:"code"`
		Message    string                 `json:"message"`
		Errors     []ErrorStructValidator `json:"errors"`
	}
)

// CustomError with Custom Error Code
type CustomError struct {
	// Code http status code
	Code int `json:"code"`

	// Message error message
	Message string `json:"message"`

	// To support more front end error handling
	// ServiceId refer to services.id in tf_auth, every service will set variable called SERVICE_ID and get the value from tf-auth
	ServiceId int `json:"serviceId,omitempty"`
	// ErrorCode should be enum, capitilize and use snake case, example: PHONE_UNVERIFIED, defined by each need of services
	ErrorCode string `json:"errorCode,omitempty"`
}

func ErrCustomError(code int, err error) *CustomError {
	return &CustomError{
		Code:    code,
		Message: err.Error(),
	}
}

func ErrCustomErrorWithInternalErrorCode(code, serviceId int, errCode string, err error) *CustomError {
	return &CustomError{
		Code:      code,
		Message:   err.Error(),
		ServiceId: serviceId,
		ErrorCode: errCode,
	}
}

func (e CustomError) Error() string {
	return e.Message
}

type CustomErrorResponse struct {
	Error CustomError `json:"error"`
}
