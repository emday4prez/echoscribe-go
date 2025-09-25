package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// lambda handler
func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// logging the request path and returning a simple message
	fmt.Printf("Processing request for path: %s\n", request.Path)

	// the response that API Gateway expects
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*", // Allow any origin
			"Access-Control-Allow-Methods": "GET,OPTIONS",
			"Access-Control-Allow-Headers": "Content-Type",
		},
		Body: `{"message": "Hello from your Go API!"}`,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
