package domain

import (
	"testing"

	"github.com/rafazx/lambda-go/internal/models"
)

type mockRepositoryPort struct {
}

func (m mockRepositoryPort) PutItem(t *models.Transfer) error {

}

func (m mockRepositoryPort) GetItem(id string) (acc *models.Account, err error) {

}

func Test_TransferDomain(t *testing.T) {
	t.Run("", func(t *testing.T) {
		NewTransferDomain()
	})
}
