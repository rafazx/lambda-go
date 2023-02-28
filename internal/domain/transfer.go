package domain

import (
	"errors"
	"log"
	"strconv"

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

func (t *TransferDomain) CreateTransfer(tran *models.Transfer, accountId string) error {
	amount := tran.Amount

	hasAmount, err := t.hasAmount(accountId, amount)

	if err != nil {
		log.Default().Println(err.Error())
		return err
	}

	if hasAmount == false {
		log.Default().Println("Account does not have enough amount" + accountId)
		return errors.New("Account does not have enough amount")
	}

	_, erro := t.repository.PutItem(tran)

	if erro != nil {
		log.Default().Println(erro.Error())
		return erro
	}

	return nil
}

func (t *TransferDomain) hasAmount(id string, tramsferAmount string) (bool, error) {
	acc, err := t.repository.GetItem(id)

	if err != nil {
		log.Default().Println(err.Error())
		return false, err
	}

	a, _ := strconv.Atoi(tramsferAmount)
	accTotal, _ := strconv.Atoi(acc.TotalAmount)

	log.Default().Println(a)
	log.Default().Println(accTotal)

	if accTotal < a {
		return false, nil
	}

	return true, nil
}
