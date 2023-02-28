package repository

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/rafazx/lambda-go/internal/models"
	"github.com/rafazx/lambda-go/internal/ports"
)

type Database struct {
	dynamoPort ports.DynamoPort
	connection dynamodbiface.DynamoDBAPI
}

func NewRepository(dynamoPort ports.DynamoPort, conn dynamodbiface.DynamoDBAPI) *Database {
	return &Database{
		dynamoPort,
		conn,
	}
}

func (db *Database) PutItem(t *models.Transfer) error {
	tr, err := dynamodbattribute.MarshalMap(t)
	if err != nil {
		return errors.New("Error in MarshalMap")
	}

	input := &dynamodb.PutItemInput{
		Item:      tr,
		TableName: aws.String(os.Getenv("DYNAMODB_TABLE_TRANSFER")),
	}

	if err != nil {
		return errors.New("Error in create PutItemInput")
	}

	_, err := db.connection.PutItem(input)

	if err != nil {
		return err
	}

	return nil
}

func (db *Database) GetItem(id string) (acc *models.Account, err error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("DYNAMODB_TABLE_ACCOUNT")),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}

	result, err := db.connection.GetItem(input)

	if err != nil {
		log.Default().Println("Got error calling GetItem:", err)
		return nil, err
	}

	if result.Item == nil {
		return nil, errors.New("Could not find '" + id + "'")
	}

	var account *models.Account

	err = dynamodbattribute.UnmarshalMap(result.Item, &account)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	return account, nil
}
