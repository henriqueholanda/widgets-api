package widgets

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/satori/go.uuid"
	"fmt"
	"github.com/henriqueholanda/widgets-api/services"
	"os"
)

func fetchAll() (Widgets, error) {
	scanInput := &dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("WIDGETS_TABLE")),
	}

	result, err := services.GetDatabaseSession().Scan(scanInput)
	if err != nil {
		return Widgets{}, err
	}
	if len(result.Items) == 0 {
		return Widgets{}, nil
	}

	widgets := Widgets{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &widgets)
	if err != nil {
		return Widgets{}, err
	}

	return widgets, nil
}

func fetchOne(id string) (Widget, error) {
	inputItem := &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("WIDGETS_TABLE")),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}

	result, err := services.GetDatabaseSession().GetItem(inputItem)
	if err != nil {
		return Widget{}, err
	}
	if result.Item == nil {
		return Widget{}, nil
	}

	widget := Widget{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &widget)
	if err != nil {
		return Widget{}, err
	}

	return widget, nil
}

func create(widget Widget) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv("WIDGETS_TABLE")),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(fmt.Sprintf("%v", uuid.NewV4())),
			},
			"name": {
				S: aws.String(widget.Name),
			},
			"color": {
				S: aws.String(widget.Color),
			},
			"price": {
				S: aws.String(widget.Price),
			},
			"melts": {
				BOOL: aws.Bool(widget.Melts),
			},
			"inventory": {
				S: aws.String(widget.Inventory),
			},
		},
	}

	_, err := services.GetDatabaseSession().PutItem(input)
	return err
}

func update(widget Widget) error {
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(os.Getenv("WIDGETS_TABLE")),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(widget.ID),
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#name": aws.String("name"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":n": {
				S: aws.String(widget.Name),
			},
			":c": {
				S: aws.String(widget.Color),
			},
			":p": {
				S: aws.String(widget.Price),
			},
			":m": {
				BOOL: aws.Bool(widget.Melts),
			},
			":i": {
				S: aws.String(widget.Inventory),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set " +
			"#name = :n, " +
			"color = :c, " +
			"price = :p, " +
			"melts = :m, " +
			"inventory = :i"),
	}

	_, err := services.GetDatabaseSession().UpdateItem(input)
	return err
}
