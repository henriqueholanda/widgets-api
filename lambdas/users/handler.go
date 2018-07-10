package users

import (
	"github.com/aws/aws-lambda-go/events"
	"fmt"
	"encoding/json"
	"github.com/henriqueholanda/widgets-api/response"
)

func HandlerGetAllUsers(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	users, err := fetchAll()
	if err != nil {
		fmt.Println(err)
		return response.InternalServerError()
	}

	usersResponse, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return response.InternalServerError()
	}

	return response.OkResponse(string(usersResponse))
}

func HandlerGetSingleUser(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := req.PathParameters["id"]

	user, err := fetchOne(id)
	if err != nil {
		fmt.Println(err)
		return response.InternalServerError()
	}

	if user.ID == "" {
		return response.NotFound()
	}

	userResponse, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return response.InternalServerError()
	}

	return response.OkResponse(string(userResponse))
}
