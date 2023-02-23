package domain

import (
	"log"

	"github.com/rafazx/lambda-go/internal/models"
	"github.com/rafazx/lambda-go/internal/ports"
)

type TransferDomain struct {
	repository ports.DynamoPort
}

func NewTransferDomain(repository ports.DynamoPort) *TransferDomain {
	return &TransferDomain{
		repository,
	}
}

func (t *TransferDomain) CreateTransfer(tran *models.Transfer) (string, error) {
	_, err := t.repository.PutItem(tran)

	if err != nil {
		log.Fatal("Deu ruimmmmmmmmmmmmm")
	}

	return "", nil
}
