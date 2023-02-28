package ports

import (
	"github.com/rafazx/lambda-go/internal/models"
)

type DynamoPort interface {
	PutItem(t *models.Transfer) error
	GetItem(id string) (acc *models.Account, err error)
}
