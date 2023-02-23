package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rafazx/lambda-go/internal/adpters/http"
	"github.com/rafazx/lambda-go/internal/adpters/repository"
	"github.com/rafazx/lambda-go/internal/domain"
	"github.com/rafazx/lambda-go/internal/instance"
	"github.com/rafazx/lambda-go/internal/ports"
)

var (
	repPort  ports.DynamoPort
	httpPort ports.HttpPort

	HttpAdapter http.HttpAdapter
)

func init() {
	conn := instance.GetConnection()
	repo := repository.NewRepository(repPort, conn)

	d := domain.NewTransferDomain(repo)
	httpPort = d
}

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return HttpAdapter.HandleHttp(ctx, req)
}

func main() {
	lambda.Start(HandleRequest)
}
