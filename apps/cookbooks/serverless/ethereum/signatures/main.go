package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, event interface{}) (string, error) {
	fmt.Printf("%v\n", event)
	return "Hello from Lambda!", nil
}

func main() {
	lambda.Start(HandleRequest)
}
