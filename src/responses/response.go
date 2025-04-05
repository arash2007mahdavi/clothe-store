package responses

import "store/src/validation"

type Response struct {
	Status        bool                          `json:"Status"`
	StatusCode    int                           `json:"StatusCode"`
	Response      any                           `json:"Response,omitempty"`
	Err           string                        `json:"Error,omitempty"`
	Validationerr *[]validation.Validationerror `json:"ValidationError,omitempty"`
}

func MakeNormalResponse(status bool, statusCode int, response any) *Response {
	return &Response{
		Status:     status,
		StatusCode: statusCode,
		Response:   response,
	}
}

func MakeResponseWithError(status bool, statusCode int, err error) *Response {
	return &Response{
		Status:     status,
		StatusCode: statusCode,
		Err:        err.Error(),
	}
}

func MakeResponseWithValidationError(status bool, statusCode int, err error) *Response {
	return &Response{
		Status:        status,
		StatusCode:    statusCode,
		Validationerr: validation.MakeValidationError(err),
	}
}
