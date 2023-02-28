package http_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rafazx/lambda-go/internal/adpters/http"
	"github.com/rafazx/lambda-go/internal/models"
)

type mockPorts struct {
}

func (m mockPorts) CreateTransfer(tran *models.Transfer, accountId string) error {
	if tran.Id != "" {
		return nil
	}

	return errors.New("Body is empty")
}

func Test_HttpAdpater(t *testing.T) {
	t.Run("Should return status 200", func(t *testing.T) {
		h := http.NewHttpAdapter(mockPorts{})

		b := `{"id": "431","amount": "100","is_processed": false}`

		e := events.APIGatewayProxyRequest{
			QueryStringParameters: map[string]string{
				"account_id": "123",
			},
			Body: b,
		}

		resp, err := h.HandleHttp(e)

		if !reflect.DeepEqual(resp.StatusCode, 200) {
			t.Fatal("Status is not 400")
		}

		if !errors.Is(err, nil) {
			t.Fatal("Error is not empty")
		}
	})

	t.Run("Should return status 400 when body is empty", func(t *testing.T) {
		h := http.NewHttpAdapter(mockPorts{})

		b := `{}`

		e := events.APIGatewayProxyRequest{
			QueryStringParameters: map[string]string{
				"account_id": "123",
			},
			Body: b,
		}

		resp, err := h.HandleHttp(e)

		if !reflect.DeepEqual(resp.StatusCode, 400) {
			t.Fatal("Status is not 400")
		}

		if !errors.Is(err, nil) {
			t.Fatal("Error is not empty")
		}

	})

}
