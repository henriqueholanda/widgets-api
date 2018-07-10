package widgets

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/henriqueholanda/widgets-api/response"
	"encoding/json"
)

func HandlerGetAllWidgets(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	widgets, err := fetchAll()
	if err != nil {
		return response.InternalServerError()
	}

	widgetsResponse, err := json.Marshal(widgets)
	if err != nil {
		return response.InternalServerError()
	}

	return response.OkResponse(string(widgetsResponse))
}

func HandlerGetSingleWidget(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := req.PathParameters["id"]

	widget, err := fetchOne(id)
	if err != nil {
		return response.InternalServerError()
	}

	if widget.ID == "" {
		return response.NotFound()
	}

	widgetResponse, err := json.Marshal(widget)
	if err != nil {
		return response.InternalServerError()
	}

	return response.OkResponse(string(widgetResponse))
}

func HandlerCreateWidget(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	widget := Widget{}
	err := json.Unmarshal([]byte(req.Body), &widget)
	if err != nil {
		return response.InternalServerError()
	}

	err = create(widget)
	if err != nil {
		return response.InternalServerError()
	}

	return response.Created()
}

func HandlerUpdateWidget(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := req.PathParameters["id"]

	widget, err := fetchOne(id)
	if err != nil {
		return response.NotFound()
	}

	err = json.Unmarshal([]byte(req.Body), &widget)
	if err != nil {
		return response.InternalServerError()
	}

	err = update(widget)
	if err != nil {
		return response.InternalServerError()
	}

	widgetResponse, err := json.Marshal(widget)
	if err != nil {
		return response.InternalServerError()
	}
	return response.OkResponse(string(widgetResponse))
}
