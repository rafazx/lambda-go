package repository

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/rafazx/lambda-go/internal/models"
)

type mockConnection struct {
	dynamodbiface.DynamoDBAPI
	getOut *dynamodb.GetItemOutput
	putOut *dynamodb.PutItemOutput
	err    error
}

func (c mockConnection) PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return c.putOut, c.err
}

func (c mockConnection) GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return c.getOut, c.err
}

type mockDatabasePort struct {
}

func (d mockDatabasePort) PutItem(t *models.Transfer) error {
	return nil
}

func (d mockDatabasePort) GetItem(id string) (acc *models.Account, err error) {
	r := &models.Account{
		Id:          "1234",
		Name:        "tester",
		TotalAmount: "30000",
	}

	return r, nil
}

func Test_Repository(t *testing.T) {
	repository := NewRepository(mockDatabasePort{}, mockConnection{
		getOut: &dynamodb.GetItemOutput{
			Item: map[string]*dynamodb.AttributeValue{},
		},
	})

	t.Run("Should call PutItem without error", func(t *testing.T) {
		var tr *models.Transfer

		err := repository.PutItem(tr)

		if !errors.Is(err, nil) {
			t.Fatal("Error is not nil")
		}
	})

	t.Run("Should return error in dynamo PutItem", func(t *testing.T) {
		var tr *models.Transfer

		e := errors.New("Error in Dynamo")

		repository := NewRepository(mockDatabasePort{}, mockConnection{
			err: e,
		})

		err := repository.PutItem(tr)

		if !errors.Is(err, e) {
			t.Fatal("Unexpected error", err)
		}
	})

	t.Run("Should call GetItem without error", func(t *testing.T) {
		_, err := repository.GetItem("1234")

		if !errors.Is(err, nil) {
			t.Fatal("Error is not nil")
		}
	})

	t.Run("Should call GetItem with error", func(t *testing.T) {

		e := errors.New("Error in Dynamo")

		repository := NewRepository(mockDatabasePort{}, mockConnection{
			err: e,
		})

		_, err := repository.GetItem("1234")

		if !errors.Is(err, e) {
			t.Fatal("Unexpected error", err)
		}
	})

}
