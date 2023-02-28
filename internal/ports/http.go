package ports

import "github.com/rafazx/lambda-go/internal/models"

type HttpPort interface {
	CreateTransfer(tran *models.Transfer, accountId string) error
}
