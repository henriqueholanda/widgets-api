package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"regexp"
	"github.com/aws/aws-lambda-go/events"
	"github.com/henriqueholanda/widgets-api/response"
	"github.com/henriqueholanda/widgets-api/lambdas/users"
	"github.com/henriqueholanda/widgets-api/lambdas/widgets"
)

const usersEndpoint			= "/users"
const widgetsEndpoint  		= "/widgets"

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if req.HTTPMethod == "GET" {
		hasUserID, _ := regexp.MatchString(usersEndpoint + "/.+", req.Path)
		if hasUserID {
			return users.HandlerGetSingleUser(req)
		}

		if req.Path == usersEndpoint {
			return users.HandlerGetAllUsers(req)
		}

		hasWidgetID, _ := regexp.MatchString(widgetsEndpoint + "/.+", req.Path)
		if hasWidgetID {
			return widgets.HandlerGetSingleWidget(req)
		}

		if req.Path == widgetsEndpoint {
			return widgets.HandlerGetAllWidgets(req)
		}

		return response.NotFound()
	}

	if req.HTTPMethod == "POST" {
		if req.Path == widgetsEndpoint {
			return widgets.HandlerCreateWidget(req)
		}

		return response.NotFound()
	}

	if req.HTTPMethod == "PUT" {
		hasWidgetID, _ := regexp.MatchString(widgetsEndpoint + "/.+", req.Path)
		if hasWidgetID {
			return widgets.HandlerUpdateWidget(req)
		}

		return response.NotFound()
	}

	return response.MethodNotAllowed()
}

func main() {
	lambda.Start(router)
}
