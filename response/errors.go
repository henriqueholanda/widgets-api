package response

import (
	"net/http"

	"github.com/henriqueholanda/widgets-api/services"
)

func MethodNotAllowed() (services.Response, error) {
	return services.Response{
		StatusCode: http.StatusMethodNotAllowed,
		Body:       http.StatusText(http.StatusMethodNotAllowed),
	}, nil
}

func NotFound() (services.Response, error) {
	return services.Response{
		StatusCode: http.StatusNotFound,
		Body:       http.StatusText(http.StatusNotFound),
	}, nil
}

func InternalServerError() (services.Response, error) {
	return services.Response{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

func Unauthorized() (services.Response, error) {
	return services.Response{
		StatusCode: http.StatusUnauthorized,
		Body:       http.StatusText(http.StatusUnauthorized),
	}, nil
}
