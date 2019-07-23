package widgets

import (
	"encoding/json"

	"github.com/henriqueholanda/widgets-api/response"
	"github.com/henriqueholanda/widgets-api/services"
)

// HandlerGetAllWidgets Returns a List of Widgets from database
// @Title All Widgets Endpoint
// @Description Returns a List of Widgets from database
// @Success 200 {array} Widgets "A list of widgets"
// @Param Authentication header string true "Authentication Token"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @router /widgets [get]
func HandlerGetAllWidgets(request services.Request) (services.Response, error) {
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

// HandlerGetSingleWidget Returns a Single widget from database
// @Title Single Widget Endpoint
// @Description Returns a Single Widget from database
// @Success 200 {json} Widget "A single widget"
// @Param id path string true "Widget ID"
// @Param Authentication header string true "Authentication Token"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @router /widgets/{id} [get]
func HandlerGetSingleWidget(request services.Request) (services.Response, error) {
	id := request.PathParameters["id"]

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

// HandlerCreateWidget Create a widget on database
// @Title Create Widget Endpoint
// @Description Create a widget on database
// @Success 201 {object} string "Created"
// @Param widget body Widget true "Widget"
// @Param Authentication header string true "Authentication Token"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @router /widgets [post]
func HandlerCreateWidget(request services.Request) (services.Response, error) {
	widget := Widget{}
	err := json.Unmarshal([]byte(request.Body), &widget)
	if err != nil {
		return response.InternalServerError()
	}

	err = create(widget)
	if err != nil {
		return response.InternalServerError()
	}

	return response.Created()
}

// HandlerUpdateWidget Update a widget on database and return itself
// @Title Update Widget Endpoint
// @Description Update a widget on database and return itself
// @Success 200 {json} Widget "A updated widget"
// @Param id path string true "Widget ID"
// @Param widget body Widget true "Widget"
// @Param Authentication header string true "Authentication Token"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @router /widgets/{id} [put]
func HandlerUpdateWidget(request services.Request) (services.Response, error) {
	id := request.PathParameters["id"]

	widget, err := fetchOne(id)
	if err != nil {
		return response.NotFound()
	}

	err = json.Unmarshal([]byte(request.Body), &widget)
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
