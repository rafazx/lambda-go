package http

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rafazx/lambda-go/internal/models"
	"github.com/rafazx/lambda-go/internal/ports"
)

type HttpAdapter struct {
	httpPort ports.HttpPort
}

func NewHttpAdapter(httpPort ports.HttpPort) *HttpAdapter {
	return &HttpAdapter{
		httpPort,
	}
}

func (h *HttpAdapter) HandleHttp(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	accountId := req.QueryStringParameters["account_id"]

	var tran *models.Transfer
	json.Unmarshal([]byte(req.Body), &tran)

	err := h.httpPort.CreateTransfer(tran, accountId)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Item Criado com sucesso",
	}, nil
}
