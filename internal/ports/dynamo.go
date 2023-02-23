package ports

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/rafazx/lambda-go/internal/models"
)

type DynamoPort interface {
	PutItem(t *models.Transfer) (response *dynamodb.PutItemOutput, err error)
}
