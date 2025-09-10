package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// HandleRequest is our lambda handler
func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// For now, we're just logging the request path and returning a simple message.
	fmt.Printf("Processing request for path: %s\n", request.Path)

	// This is the response that API Gateway expects
	return events.APIGatewayProxyResponse{
		Body:       `{"message": "Hello from your Go API!"}`,
		StatusCode: 200,
	}, nil
}

func main() {
	// This starts the handler
	lambda.Start(HandleRequest)
}
