package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rafazx/lambda-go/internal/adpters/http"
	"github.com/rafazx/lambda-go/internal/adpters/repository"
	"github.com/rafazx/lambda-go/internal/domain"
	"github.com/rafazx/lambda-go/internal/ports"
)

var (
	repPort ports.DynamoPort

	HttpAdapter *http.HttpAdapter
)

func init() {
	log.Default().Println("Factory")

	conn := repository.GetConnection()
	repo := repository.NewRepository(repPort, conn)

	d := domain.NewTransferDomain(repo)

	HttpAdapter = http.NewHttpAdapter(d)
}

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return HttpAdapter.HandleHttp(req)
}

func main() {
	lambda.Start(HandleRequest)
}
