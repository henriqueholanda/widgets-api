package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dgrijalva/jwt-go"
	"github.com/henriqueholanda/widgets-api/lambdas/widgets"
	"github.com/henriqueholanda/widgets-api/response"
	"github.com/henriqueholanda/widgets-api/services"
	"os"
	"regexp"
	"strings"
)

const widgetsEndpoint = "/widgets"

func hasValidToken(request services.Request) bool {
	authHeader := request.Headers["Authorization"]
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

func router(request services.Request) (services.Response, error) {

	if !hasValidToken(request) {
		return response.Unauthorized()
	}

	if !isWidgetsEndpoint(request) {
		return response.NotFound()
	}

	if request.HTTPMethod == "GET" {
		if hasWidgetID(request) {
			return widgets.HandlerGetSingleWidget(request)
		}

		return widgets.HandlerGetAllWidgets(request)
	}

	if request.HTTPMethod == "POST" {
		return widgets.HandlerCreateWidget(request)
	}

	if request.HTTPMethod == "PUT" {
		if hasWidgetID(request) {
			return widgets.HandlerUpdateWidget(request)
		}

		return response.NotFound()
	}

	return response.MethodNotAllowed()
}

func isWidgetsEndpoint(request services.Request) bool {
	isWidgetsEndpoint, _ := regexp.MatchString(widgetsEndpoint, request.Path)
	return isWidgetsEndpoint
}

func hasWidgetID(request services.Request) bool {
	hasWidgetID, _ := regexp.MatchString(widgetsEndpoint + "/.+", request.Path)
	return hasWidgetID
}

// @APIVersion 2.0.0
// @APITitle Widgets API
// @APIDescription An API to manage Widgets
// @License Copyright
func main() {
	lambda.Start(router)
}
