package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, secretID string, msg any) (string, error) {
	return "Hello from Lambda!", nil
}

func main() {
	lambda.Start(HandleRequest)
}
