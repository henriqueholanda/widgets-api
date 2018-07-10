package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"regexp"
	"github.com/aws/aws-lambda-go/events"
	"github.com/dgrijalva/jwt-go"
	"github.com/henriqueholanda/widgets-api/response"
	"github.com/henriqueholanda/widgets-api/lambdas/users"
	"github.com/henriqueholanda/widgets-api/lambdas/widgets"
	"strings"
	"os"
	"fmt"
)

const usersEndpoint			= "/users"
const widgetsEndpoint  		= "/widgets"

func hasValidToken(req events.APIGatewayProxyRequest) (bool) {
	authHeader := req.Headers["Authorization"]
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_TOKEN")), nil
	})

	if err == nil && token.Valid {
		return true
	}
	if ve, ok := err.(*jwt.ValidationError); ok {
		fmt.Println("Unauthorized: " + ve.Error())
		return false
	}

	return false
}

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if !hasValidToken(req) {
		return response.Unauthorized()
	}

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

// @APIVersion 2.0.0
// @APITitle Widgets API
// @APIDescription An API to manage Widgets
// @License Copyright
func main() {
	lambda.Start(router)
}
