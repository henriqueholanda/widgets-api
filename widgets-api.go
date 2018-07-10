package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"regexp"
	"github.com/aws/aws-lambda-go/events"
	"github.com/henriqueholanda/widgets-api/response"
	"github.com/henriqueholanda/widgets-api/lambdas/users"
)

const usersEndpoint			= "/users"

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if req.HTTPMethod == "GET" {
		hasUserID, _ := regexp.MatchString(usersEndpoint + "/.+", req.Path)
		if hasUserID {
			return users.HandlerGetSingleUser(req)
		}

		if req.Path == usersEndpoint {
			return users.HandlerGetAllUsers(req)
		}

		return response.NotFound()
	}

	return response.MethodNotAllowed()
}

func main() {
	lambda.Start(router)
}
