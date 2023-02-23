package instance

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetConnection() *dynamodb.DynamoDB {
	var sess *session.Session

	if os.Getenv("AWS_STAGE") == "local" {
		sess = session.Must(session.NewSession(&aws.Config{
			Endpoint: aws.String("http://localhost:4566"),
			Region:   aws.String("us-east-1"),
		}))
	} else {
		sess = session.Must(session.NewSession())
	}

	return dynamodb.New(sess)
}
