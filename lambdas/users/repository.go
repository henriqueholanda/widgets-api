package users

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/henriqueholanda/widgets-api/services"
	"os"
)

func fetchAll() (Users, error) {
	scanInput := &dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("USERS_TABLE")),
	}

	result, err := services.GetDatabaseSession().Scan(scanInput)
	if err != nil {
		return Users{}, err
	}
	if len(result.Items) == 0 {
		return Users{}, nil
	}

	users := Users{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &users)
	if err != nil {
		return Users{}, err
	}

	return users, nil
}

func fetchOne(id string) (User, error) {
	inputItem := &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("USERS_TABLE")),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}

	result, err := services.GetDatabaseSession().GetItem(inputItem)
	if err != nil {
		return User{}, err
	}
	if result.Item == nil {
		return User{}, nil
	}

	user := User{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
