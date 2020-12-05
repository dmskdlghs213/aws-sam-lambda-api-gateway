package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	names := request.QueryStringParameters

	var name string
	for _, val := range names {
		name = val
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %v", name + " "),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
