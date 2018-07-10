package users

import (
	"github.com/aws/aws-lambda-go/events"
	"fmt"
	"encoding/json"
	"github.com/henriqueholanda/widgets-api/response"
)

// HandlerGetAllUsers Returns a List of Users from database
// @Title All Users Endpoint
// @Description Returns a List of Users from database
// @Success 200 {array} Users "A list of users"
// @Param Authentication header string true "Authentication Token"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @router /users [get]
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

// HandlerGetSingleUser Returns a Single user from database
// @Title Single User Endpoint
// @Description Returns a Single User from database
// @Success 200 {json} User "A single user"
// @Param id path string true "User ID"
// @Param Authentication header string true "Authentication Token"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @router /users/{id} [get]
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
