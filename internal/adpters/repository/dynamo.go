package repository

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/rafazx/lambda-go/internal/models"
	"github.com/rafazx/lambda-go/internal/ports"
)

type Database struct {
	dynamoPort ports.DynamoPort
	connection *dynamodb.DynamoDB
}

func NewRepository(dynamoPort ports.DynamoPort, conn *dynamodb.DynamoDB) *Database {
	return &Database{
		dynamoPort,
		conn,
	}
}

func (db *Database) PutItem(t *models.Transfer) (response *dynamodb.PutItemOutput, err error) {
	tr, err := dynamodbattribute.MarshalMap(t)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      tr,
		TableName: aws.String(os.Getenv("DYNAMODB_TABLE")),
	}

	if err != nil {
		return nil, err
	}

	return db.connection.PutItem(input)
}
