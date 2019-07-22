package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

func GetDatabaseSession() *dynamodb.DynamoDB {
	awsSession, _ := session.NewSession()

	connection := dynamodb.New(
		awsSession,
		aws.NewConfig().WithRegion(os.Getenv("REGION_AWS")),
	)
	return connection
}
