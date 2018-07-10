package response

import (
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

func OkResponse(body string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       body,
	}, nil
}

func Created() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       http.StatusText(http.StatusCreated),
	}, nil
}
