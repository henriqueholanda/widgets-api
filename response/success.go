package response

import (
	"net/http"

	"github.com/henriqueholanda/widgets-api/services"
)

func OkResponse(body string) (services.Response, error) {
	return services.Response{
		StatusCode: http.StatusOK,
		Body:       body,
	}, nil
}

func Created() (services.Response, error) {
	return services.Response{
		StatusCode: http.StatusCreated,
		Body:       http.StatusText(http.StatusCreated),
	}, nil
}
